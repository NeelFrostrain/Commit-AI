# Supported File Types

Commit-AI intelligently filters files to analyze only text-based code changes and ignores binary files that cannot be meaningfully analyzed by AI.

## Text Files (Analyzed by AI)

The following file types are analyzed by the AI to generate commit messages:

### Programming Languages
- `.go`, `.py`, `.js`, `.ts`, `.jsx`, `.tsx`, `.java`, `.c`, `.cpp`, `.h`, `.hpp`
- `.cs`, `.php`, `.rb`, `.rs`, `.swift`, `.kt`, `.scala`, `.clj`
- `.sh`, `.bash`, `.ps1`, `.bat`, `.cmd`

### Web Development
- `.html`, `.htm`, `.css`, `.scss`, `.sass`, `.less`
- `.vue`, `.svelte`, `.astro`
- `.json`, `.xml`, `.yaml`, `.yml`, `.toml`

### Configuration & Documentation
- `.md`, `.txt`, `.rst`, `.adoc`
- `.env`, `.env.example`, `.gitignore`, `.gitattributes`
- `Makefile`, `Dockerfile`, `docker-compose.yml`
- `.editorconfig`, `.prettierrc`, `.eslintrc`

### Game Development Scripts
- `.cs` (Unity C# scripts)
- `.shader`, `.cginc`, `.hlsl` (Shader files)
- `.uss`, `.uxml` (Unity UI Toolkit)
- `.h`, `.cpp` (Unreal Engine C++)

## Binary Files (Automatically Excluded)

The following file types are automatically excluded from AI analysis:

### Executables & Libraries
- `.exe`, `.dll`, `.so`, `.dylib`, `.a`, `.lib`, `.o`

### Archives
- `.zip`, `.tar`, `.gz`, `.7z`, `.rar`

### Images
- `.png`, `.jpg`, `.jpeg`, `.gif`, `.ico`, `.bmp`, `.tiff`, `.webp`
- `.psd`, `.ai` (Adobe formats)

### Unity Engine Files
- `.unity` (Scene files)
- `.prefab` (Prefab assets)
- `.asset` (Generic Unity assets)
- `.mat` (Materials)
- `.controller`, `.anim`, `.overrideController` (Animation)
- `.mask`, `.cubemap`, `.flare` (Rendering)
- `.physicMaterial`, `.physicsMaterial2D` (Physics)
- `.guiskin`, `.fontsettings` (UI)
- `.spriteatlas`, `.terrainlayer` (Sprites & Terrain)

### 3D Models & Assets
- `.fbx`, `.obj`, `.dae`, `.3ds`, `.dxf` (3D models)
- `.max` (3ds Max)
- `.blend` (Blender)
- `.mb`, `.ma` (Maya)

### Unreal Engine Files
- `.uasset` (Unreal assets)
- `.umap` (Unreal maps)
- `.upk`, `.udk`, `.u` (Unreal packages)

### Audio Files
- `.mp3`, `.wav`, `.ogg`, `.flac`, `.aac`, `.wma`, `.m4a`

### Video Files
- `.mp4`, `.avi`, `.mov`, `.wmv`, `.flv`, `.mkv`, `.webm`

### Textures & Materials
- `.dds`, `.tga`, `.exr`, `.hdr`

### Font Files
- `.ttf`, `.otf`, `.woff`, `.woff2`, `.eot`

### Other Binary Formats
- `.pdf` (Documents)
- `.db`, `.sqlite`, `.sqlite3` (Databases)
- `.dat`, `.bin`, `.pak`, `.cache` (Generic binary)

## Why Exclude Binary Files?

Binary files are excluded because:

1. **Unreadable Content**: Binary files contain compiled code or encoded data that appears as random characters when read as text
2. **AI Confusion**: Analyzing binary content leads to nonsensical commit messages
3. **Performance**: Binary files are often large and waste API tokens
4. **Irrelevant Changes**: Binary file changes are typically build artifacts, not meaningful code changes

## How It Works

Commit-AI uses two mechanisms to exclude binary files:

1. **`.gitattributes`**: Marks files as binary so Git doesn't attempt to diff them
2. **Pattern Matching**: The `internal/git/diff.go` module filters staged files by extension

## Adding Custom Patterns

To exclude additional file types, edit:

1. **`.gitattributes`**: Add pattern with `binary` attribute
   ```
   *.custom binary
   ```

2. **`internal/git/diff.go`**: Add pattern to `binaryPatterns` array
   ```go
   binaryPatterns := []string{
       // ... existing patterns
       "*.custom",
   }
   ```

## Game Development Best Practices

### Unity Projects
- Commit `.cs` scripts, `.shader` files, and `.json` configs
- Exclude `.unity`, `.prefab`, `.asset` files (use Unity's YAML mode if you need to track these)
- Use `.meta` files for asset tracking (text-based)

### Unreal Engine Projects
- Commit `.h`, `.cpp` source files and `.ini` configs
- Exclude `.uasset`, `.umap` binary assets
- Use source control integration for large binary assets

### General Game Development
- Commit source code and configuration files
- Use Git LFS for large binary assets (models, textures, audio)
- Keep binary assets in separate repositories or asset management systems

## Troubleshooting

### AI Analyzing Wrong Files?

Check if the file extension is in the exclude list:
```bash
git check-attr binary path/to/file.ext
```

If it shows `binary: set`, the file is properly marked.

### Need to Analyze a Binary File?

This is not recommended, but you can temporarily remove the pattern from:
1. `.gitattributes`
2. `internal/git/diff.go` binaryPatterns array

## Related Documentation

- [Build Guide](BUILD_GUIDE.md)
- [Contributing Guide](CONTRIBUTING.md)
- [Changelog](CHANGELOG.md)
