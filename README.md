# ✏️ Notegen

This is a simple utility that goes through your markdown notes and converts them to plain html files, to be served from your site or a static hosting service like GitHub Pages, Netlify, etc.

### Installation

You can get the appropriate executable for your system by visiting the [Releases](https://github.com/justsharan/notegen/releases) page.

On MacOS, you can install notegen from my homebrew tap:

```
brew tap justsharan/tap && brew install notegen
# OR
brew install justsharan/tap/notegen
```

### Usage

```
Usage of notegen:
  -latex
    Whether to render LaTeX equations
  -out string
    The output directory to put files in (default ".")
  -src string
    The source directory to read files from (default ".")
```

### Example

I originally designed notegen to be used for [my notes](https://github.com/justsharan/uni-notes), but it worked well enough that I thought it would be helpful to tidy it up and make it into an easily usable CLI tool. You can see the end result [here](https://notes.justsharan.xyz/BIOL205/cnidarians.html), for example.
