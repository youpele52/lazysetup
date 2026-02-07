# Tools Panel Empty Bug Fix

**Date:** 2026-02-01 | **Status:** ✅ RESOLVED

## Problem
Panel 3 (Tools) rendered empty despite containing 27 tools. Navigation (arrows, g, G) didn't work.

## Root Cause: Double-Scrolling Conflict

### The Bug
```go
v.SetOrigin(0, offset)              // Tell gocui to scroll by offset
startIdx := offset
endIdx := offset + visibleCount
for i := startIdx; i < endIdx; i++ { // Render only visible slice
    fmt.Fprintf(v, "%s\n", tool)
}
```

### Why It Failed
```
Buffer (what we write):           Display (what shows):
┌─────────────┐                   ┌─────────────┐
│ git    (0)  │ ◄─ Render here    │             │ ◄─ SetOrigin scrolls HERE
│ docker (1)  │                   │             │    (rows 10-20)
│ ...         │                   │             │
│ fzf    (10) │                   │   EMPTY!    │
├─────────────┤ ◄─ offset=10      │             │
│ (nothing)   │                   │             │
│ (nothing)   │                   │             │
└─────────────┘                   └─────────────┘
```

**Problem:** Wrote to rows 0-10, then SetOrigin scrolled view to show rows 10-20 (empty!).

### Additional Issue: Wrong Cursor Position
```go
v.SetCursor(0, cursor)  // cursor=26 (absolute)
```
SetCursor expects **relative** position (0 to visibleCount-1), not absolute (0-27).

## Debug Process

Added logging:
```go
fmt.Printf("DEBUG: startY=%d, height=%d, visibleCount=%d\n", ...)
fmt.Printf("DEBUG: Rendering tool %d: %s\n", i, tool)
```

Output showed:
- ✅ Panel coordinates correct (startY=12, height=13)
- ✅ Loop executing (rendering all 27 tools)
- ❌ Display still empty → **Scrolling bug confirmed**

## Solution

### Fix: Render ALL + Use SetOrigin Correctly

```go
// Calculate offset to keep cursor visible
if cursor < offset {
    offset = cursor
} else if cursor >= offset + visibleCount {
    offset = cursor - visibleCount + 1
}

v.SetOrigin(0, offset)  // Scroll view

// Render ALL tools (not just slice)
for i := 0; i < len(state.Tools); i++ {
    fmt.Fprintf(v, "%s %s\n", marker, tool)
}

// Cursor RELATIVE to visible area
v.SetCursor(0, cursor - offset)  // ◄─ KEY FIX
```

### How It Works
```
Tools Array (27):      Buffer (ALL):         Display (offset=10):
┌──────────┐           ┌──────────┐          ┌──────────┐
│ 0: git   │────────>  │ git      │          │          │
│ 1: docker│           │ docker   │          │          │
│ ...      │           │ ...      │   ┌─────>│ fd       │ ◄─ offset=10
│ 10: fd   │           │ fd       │───┘      │ bat      │
│ 11: bat  │           │ bat      │          │ jq       │
│ ...      │           │ ...      │          │ ...      │
│ 20: tree │           │ tree     │          │ tree     │ ◄─ cursor=20
│ ...      │           │ ...      │          │          │    (relative=10)
│ 26: sql  │           │ lazysql  │          └──────────┘
└──────────┘           └──────────┘
  Render ALL         SetOrigin(10)         Shows rows 10-20
```

## Before vs After

| Aspect | Before | After |
|--------|--------|-------|
| Display | ❌ Empty | ✅ Shows 11 tools at a time |
| Scrolling | ❌ Broken | ✅ Works (arrows, g, G) |
| Navigation | ❌ Stuck at top | ✅ All 27 tools accessible |

## Key Learnings

### 1. Don't Mix Scrolling Methods
```go
// ❌ BAD
v.SetOrigin(offset)
render items[offset:end]  // Double-scroll!

// ✅ GOOD
v.SetOrigin(offset)
render ALL items          // Let SetOrigin handle visibility
v.SetCursor(relative)     // Cursor relative to visible area
```

### 2. SetCursor Coordinates

| Method | Coordinate System |
|--------|-------------------|
| `SetOrigin(x, y)` | Absolute (buffer offset) |
| `SetCursor(x, y)` | **Relative** (to visible area) |

### 3. Debug Early
Without logging, we assumed:
- Panel not created → **Wrong** (it was)
- Tools empty → **Wrong** (27 items present)
- Loop not executing → **Wrong** (it was)

Logging revealed the real issue: scrolling conflict.

## Files Modified
- `pkg/ui/layout_panels.go` - Fixed `renderToolsPanel()` (~30 lines)
- `pkg/ui/layout_multipanel.go` - Panel height calculation (~10 lines)

## Testing
- [x] Shows tools on startup
- [x] Arrow ↓/↑ scrolls through all 27 tools
- [x] `g` jumps to first, `G` jumps to last
- [x] Works in small (80x24) and large (120x50) terminals

**Related Docs:** `PANEL_LAYOUT_FIX_PLAN.md`, `SCROLLBAR_ADDITION.md`
