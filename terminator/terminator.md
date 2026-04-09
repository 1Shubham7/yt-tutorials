# 💻 Terminator Terminal

## Installation

```bash
sudo apt install terminator        # Debian / Ubuntu / Mint
sudo yum install terminator        # RHEL / CentOS / Fedora
sudo pacman -S terminator          # Arch Linux
sudo zypper install terminator     # OpenSUSE
sudo emerge -a x11-terms/terminator # Gentoo
```

---

## Splitting & Layout

| Shortcut | Action |
|---|---|
| `Ctrl + Shift + E` | Split terminal **vertically** (side by side) |
| `Ctrl + Shift + O` | Split terminal **horizontally** (top/bottom) |
| `Ctrl + Shift + W` | Close current terminal pane |
| `Ctrl + Shift + Q` | Quit Terminator entirely |
| `Alt + L` | Open layout launcher |

---

## Navigation Between Panes

| Shortcut | Action |
|---|---|
| `Alt + ↑` | Move focus to terminal **above** |
| `Alt + ↓` | Move focus to terminal **below** |
| `Alt + ←` | Move focus to terminal on the **left** |
| `Alt + →` | Move focus to terminal on the **right** |
| `Ctrl + Shift + N` | Focus **next** terminal |
| `Ctrl + Shift + P` | Focus **previous** terminal |

---

## Tabs

| Shortcut | Action |
|---|---|
| `Ctrl + Shift + T` | Open a **new tab** |
| `Ctrl + PageDown` | Switch to **next** tab |
| `Ctrl + PageUp` | Switch to **previous** tab |
| `Ctrl + Alt + A` | Rename current **tab** title |

---

## Focus, Zoom & Fullscreen

| Shortcut | Action |
|---|---|
| `F11` | Toggle **fullscreen** window |
| `Ctrl + Shift + X` | **Maximize** current terminal — hides all other panes, fills the Terminator window. Font size stays the same. Toggle again to restore all panes. |
| `Ctrl + Shift + Z` | **Zoom** into current terminal — like maximize, but also **scales up the font and content**. Great for demos or presentations. Toggle again to restore. |

> **Maximize vs Zoom:**  
> - `Ctrl+Shift+X` = expand current pane silently (layout hidden, same font)  
> - `Ctrl+Shift+Z` = expand current pane + scale up text visually (presentation mode)

---

## Broadcast & Grouping (Power Feature!)

Send the same input to multiple terminals simultaneously.

| Shortcut | Action |
|---|---|
| `Super + G` | **Group all** terminals — input goes to all |
| `Super + Shift + G` | **Remove grouping** from all terminals |
| `Super + T` | Group all terminals **in current tab** only |
| `Super + Shift + T` | Remove grouping from current tab |
| `Alt + A` | **Broadcast to all** terminals |
| `Alt + G` | Broadcast to **grouped** terminals only |
| `Alt + O` | Turn **broadcasting off** |

> Incredibly useful for running the same command on multiple servers at once!

---

## Copy, Paste & Search

| Shortcut | Action |
|---|---|
| `Ctrl + Shift + C` | **Copy** selected text to clipboard |
| `Ctrl + Shift + V` | **Paste** from clipboard |
| `Ctrl + Shift + F` | **Search** within terminal scrollback |
| `Ctrl + Shift + S` | Toggle **scrollbar** visibility |

---

## Font Size

| Shortcut | Action |
|---|---|
| `Ctrl + +` | **Increase** font size |
| `Ctrl + -` | **Decrease** font size |
| `Ctrl + 0` | **Reset** font size to default |

---

## Reset & Clear

| Shortcut | Action |
|---|---|
| `Ctrl + Shift + R` | **Reset** terminal state (fixes misbehaving terminals) |
| `Ctrl + Shift + G` | Reset terminal state **and clear** the window |

---

## Renaming

| Shortcut | Action |
|---|---|
| `Ctrl + Alt + W` | Rename **window** title |
| `Ctrl + Alt + A` | Rename **tab** title |
| `Ctrl + Alt + X` | Rename **terminal** (pane) title |

---

## Drag & Drop

You can **rearrange terminal panes** using drag and drop:

- Click and hold on a terminal's **titlebar** (that colored strip at the top of a pane)
- Drag it to a new position in the layout
- The target zone will be highlighted before you drop

**Alternative drag method:**
1. Hold `Ctrl`, then click and hold the **right mouse button**
2. Release `Ctrl`
3. Now drag the terminal wherever you want

---

## Launch Options (CLI Flags)

Run these when starting Terminator from the command line:

| Command | What it does |
|---|---|
| `terminator -m` | Start **maximized** |
| `terminator -b` | Start **borderless** (no window decorations) |
| `terminator -H` | **Hide** window on startup (toggle with `Ctrl+Shift+Alt+A`) |
| `terminator -l <name>` | Launch with a **saved layout** |
| `terminator -e <command>` | Run a specific **command** on startup instead of your shell |
| `terminator --new-tab` | Open a new tab in an **already running** Terminator instance |
| `terminator --toggle-visibility` | Toggle visibility of a running Terminator (Wayland) |
| `terminator -d` | Enable **debug** output |

---

## Saving & Loading Layouts

1. Arrange your terminals the way you want
2. Go to **Right-click → Preferences → Layouts**
3. Click **Add** and name your layout
4. Save it, then launch it anytime with:

```bash
terminator -l my-layout-name
```

---

## Profiles & Customization

Access via **Right-click → Preferences**:

| Section | What you can change |
|---|---|
| **Profiles** | Colors, fonts, cursor shape, scrollback lines |
| **Keybindings** | Remap ANY shortcut to your preference |
| **Layouts** | Save and manage terminal arrangements |
| **Plugins** | Enable/disable extensions |

> All shortcuts in this cheat sheet are the **defaults** — every single one can be rebound in Preferences → Keybindings.

---

## Plugins

Terminator supports plugins that extend its functionality. Access via **Right-click → Preferences → Plugins**. Common plugins include:

- **ActivityWatch** — Notifies you when a terminal becomes active/quiet
- **Logger** — Logs terminal output to a file
- **CustomCommandsMenu** — Add your own commands to the right-click menu
- **LaunchpadBugURLHandler** / **LaunchpadCodeURLHandler** — Clickable Launchpad URLs

Thanks for reading : )
