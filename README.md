# 📺 One Pace Library

### Automated Episode Manager for One Pace (Go + SvelteKit)

A modern **Sonarr-style episode manager** designed specifically for **One Pace**.  
Tracks episodes, monitors new releases, downloads automatically through qBittorrent, and keeps your media library organized for Jellyfin.

---

<div align="center">
      <img  src="https://github.com/jasanpreetn9/onepace-library/blob/main/public/view.png?raw=true" >

  </div>

---

## ✨ Features

### 🔍 Metadata Integration

- Fetches metadata from the official One Pace Episode Guide
- Automatically loads:
  - Arcs
  - Episode titles
  - Descriptions
  - CRC32 checksums
  - Release dates
  - File URLs
  - Audio & subtitle languages
  - Resolution information

### 📁 Library Scanner

- Scans your media library for existing One Pace episodes
- Supports both filename formats:
  - Download format:  
    `[One Pace][1058-1059] Egghead 01 [1080p][CA3F14A8].mkv`
  - Library format:  
    `S36E01 - New Emperors [CA3F14A8].mkv`
- Automatically imports episodes into `library.json`
- Generates Jellyfin-compatible NFO files
- Handles versioning via CRC32

### 📥 Download Automation

- Add episodes to **qBittorrent** with one click
- Automatically fetches the latest version
- Planned features:
  - Upgradable episode detection
  - Missing episode download queue
  - Failed download detection

### 🌐 Web UI (SvelteKit)

A clean, dark-themed UI inspired by **Sonarr**:

- Sidebar with arcs
- Episode browser per arc
- Status badges (imported / missing / queued / failed)
- Monitor/unmonitor episodes
- Responsive modern design (vanilla CSS)

### ⚙️ Backend (Go)

- Fast, lightweight REST API
- Hot reload with Air
- Persistent library & config files
- Structured project layout

---

## 🏗 Project Structure

```
onepace-library/
│
├── backend/
│   ├── cmd/server/           # Main API server
│   ├── internal/
│   │   ├── api/              # HTTP handlers
│   │   ├── config/           # Config handling
│   │   ├── library/          # Library.json store
│   │   ├── metadata/         # Metadata fetch & cache
│   │   ├── nfo/              # NFO generator
│   │   ├── qbittorrent/      # qBittorrent API client
│   │   └── scanner/          # Library & downloads scanner
│   ├── data/
│   │   ├── config.yml        # User config
│   │   └── library.json      # Library state
│   ├── go.mod
│   ├── go.sum
│   └── .air.toml             # Hot reload configuration
│
└── frontend/
    ├── src/
    │   ├── lib/
    │   │   ├── api.ts        # Frontend API wrapper
    │   │   ├── stores.ts     # Global Svelte stores
    │   │   ├── types.ts      # Shared types
    │   │   ├── app.css       # Global theme
    │   ├── routes/
    │   │   ├── +layout.svelte
    │   │   └── +page.svelte
    ├── package.json
    └── vite.config.js
```

---

## ⚙️ Configuration

```yaml
backend/data/config.yml:

port: "8989"

libraryPath: "./media"
downloadsPath: "./downloads"

metadata:
  episodesURL: "https://raw.githubusercontent.com/.../episodes.json"
  arcsURL: "https://raw.githubusercontent.com/.../arcs.json"

qbittorrent:
  host: "http://127.0.0.1:8080"
  username: "admin"
  password: "adminadmin"
```

---

## 🚀 Running the Project

### Backend (Go)

Install Air:

```bash
go install github.com/air-verse/air@latest
export PATH="$PATH:$(go env GOPATH)/bin"

Run:

cd backend
air

Or without hot reload:

go run ./cmd/server
```

---

### Frontend (SvelteKit)

```bash
cd frontend
npm install
npm run dev
```

---

## 📡 API Endpoints

### Metadata & Library

```
GET  /api/episodes/all
POST /api/scan/library
POST /api/scan/downloads
```

### Downloads

```
POST /api/download/add
POST /api/episodes/monitor
```

---

## 🔮 Roadmap / Planned Features

### Library & Metadata

- Upgradable episode detection
- Missing episode list
- Arc progress tracking
- Episode search & filters
- Version history

### Downloads

- qBittorrent queue viewer
- Retry failed downloads
- Manual search modal

### UI Enhancements

- Episode detail modal
- Settings page
- Activity log
- Light theme

### Automation

- Auto-refresh metadata
- Auto-download monitored episodes
- Auto-clean old versions

---

## 🤝 Contributing

Pull requests welcome.  
Clean code, structured modules, easy to extend.

---

## 📄 License

GNU GPLv3 — open and free to modify.

---

## ❤️ Acknowledgements

Metadata sources from `jasanpreetn9/onepace-metadata`.