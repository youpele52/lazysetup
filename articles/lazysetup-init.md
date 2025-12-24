# Lazysetup: Because Tired of Repetitive Server Setup


December 24, 2025

3 min read

There's this thing about working with multiple Linux servers. You get a fresh server from Hetzner, DigitalOcean, or wherever, and then you realize: oh right, I need git. Then you remember: oh right, docker too. And then: lazygit would be nice. And maybe lazydocker while I'm at it.

Every. Single. Time.

I found myself going through this ritual repeatedly. SSH into a new server, search for the apt install commands, wait for the downloads, repeat. It wasn't painful, just... repetitive. The kind of repetition that makes you think there has to be a better way.

So I built lazysetup. It's a terminal UI tool that lets you install, update, and uninstall development tools across multiple Linux servers with a consistent interface. Because I got tired of writing the same installation commands over and over again.

The idea came from the fluidity of tools like lazygit, lazydocker, and opencode—terminal UIs that make complex interactions feel natural. There's something elegant about doing things in the terminal without constantly context-switching between the terminal and a web interface.

The beauty of it is that you can select your package manager—Homebrew, APT, YUM, Curl, even Scoop or Chocolatey for Windows—then choose which tools you want to install from a list. Watch the progress in real-time. Done.

## Why This Matters

Look, I get it. Installing git and docker isn't hard. You can do it in a few commands. But when you're managing multiple servers or setting up fresh environments regularly, those few commands add up. You start thinking about automation, but writing scripts for every scenario is overkill for most people.

Lazysetup gives you a middle ground. Not as complex as Ansible or Chef, not as manual as running individual commands. Just a simple interface that works consistently across different platforms.

It handles multiple actions beyond just installation—updates and uninstallations too. And it's all in a single binary you can curl down or install via Go. No complex dependencies or configuration files unless you want them.

## How It Works

Start it up:

```bash
./lazysetup
```

You'll see a multi-panel interface. Navigate between panels with Tab or number keys. Select your package manager. Choose your action—install, update, or uninstall. Select tools. Watch the progress.

The tools it supports right now: git, docker, lazygit, lazydocker. Basic stuff, but the stuff you actually need on most servers.

```
┌─────────────────────────────────────────────────┐
│ [1]-Package Manager  [2]-Action  [3]-Tools       │
├──────────────┬──────────────┬──────────────────┤
│ Methods      │ Action       │ Tools            │
│              │              │                  │
│ ○ Homebrew   │ ● Install    │ ☑ git            │
│ ○ APT        │ ○ Update     │ ☑ docker         │
│ ○ Curl       │ ○ Uninstall  │ ☐ lazygit        │
│ ○ YUM        │              │ ☐ lazydocker     │
└──────────────┴──────────────┴──────────────────┘
```

## Installation

Install it via curl—the simplest way:

```bash
curl -fsSL https://github.com/youpele52/lazysetup/releases/download/v0.0.1/install.sh | bash
```

Or if you prefer Go:

```bash
go install github.com/youpele52/lazysetup@latest
```

## What's Coming

The roadmap includes more tools—Node.js, Python, and other developer essentials. AI-powered error resolution for when installations fail. Team configuration sharing via YAML files. The typical things you'd expect from a tool like this.

But right now, it does exactly what I built it for: installs git and docker on every Linux server I work with, without me having to remember the exact package names or commands for each distribution.

Sometimes the best tools are the ones that solve a specific pain point, even if that pain point seems small.

> Keep creating and keep building.

[![arrow](/arrow-up-light.svg)](#top-of-the-world)

[Home](/)[Poetry](/poetry)[Art](/art)[Notes](/notes)[Projects](/projects)

[GitHub](https://github.com/youpele52/)[LinkedIn](https://www.linkedin.com/in/youpele52/)[Twitter](https://twitter.com/youpele52)
