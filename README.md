# moonlight-embedded-tui

Moonlight Embedded TUI (MET) is a small application to create a terminal UI (TUI) using [Bubble Tea](https://github.com/charmbracelet/bubbletea/) and [Bubbles](https://github.com/charmbracelet/bubbles/) Go frameworks to support [Moonlight Embedded](https://github.com/moonlight-stream/moonlight-embedded) installs.

Gracefully executes generated moonlight command for selected application, pausing MET while the stream is initiated, resuming when the stream ends, cleanly providing a persistent TUI. This helps cut down on repeated terminal typed commands (especially useful if you shut down your raspberry pi from a power switch where you lose terminal history or typing commands is difficult due to how you're interfacing with your RPi) and provides a clean interface.

## Demo

https://github.com/user-attachments/assets/6067b381-e307-4f4d-8a3a-7eeae4eb7ee1

# Todo list
 * Allow picking between multiple hosts
 * Populate list of applications via moonlight, per host
 * Allow setting default resolution and bitrate
 * Allow picking different resolution and bitrate from current default
