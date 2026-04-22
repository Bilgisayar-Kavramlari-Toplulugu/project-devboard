# DevBoard

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VS Code](https://code.visualstudio.com/) + [Vue (Official)](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Recommended Browser Setup

- Chromium-based browsers (Chrome, Edge, Brave, etc.):
  - [Vue.js devtools](https://chromewebstore.google.com/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd)
  - [Turn on Custom Object Formatter in Chrome DevTools](http://bit.ly/object-formatters)
- Firefox:
  - [Vue.js devtools](https://addons.mozilla.org/en-US/firefox/addon/vue-js-devtools/)
  - [Turn on Custom Object Formatter in Firefox DevTools](https://fxdx.dev/firefox-devtools-custom-object-formatters/)

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

## Docker Setup

### Build Locally

```sh
docker build -t devboard-frontend:0.0.1 .
```

### Run Locally

```sh
docker run --rm -p 5173:5173 devboard-frontend:0.0.1
```

Then open: [http://localhost:5173](http://localhost:5173)

### Custom Port

To run on a different port (e.g., 8080):

```sh
docker run --rm -e PORT=8080 -p 8080:8080 devboard-frontend:0.0.1
```

### Security Features

The Docker image includes:
- **Multi-stage build**: Removes build dependencies from final image, reducing image size and attack surface
- **Non-root user**: Application runs as `nodejs` user (UID 1001), not root
- **Pinned versions**: Node 20.20.2-alpine for consistent, secure builds
- **Optimized install**: Uses `npm ci` instead of `npm install` for reliable, reproducible installs
- **.dockerignore**: Excludes unnecessary files (node_modules, git, tests, etc.) from build context

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Run End-to-End Tests with [Playwright](https://playwright.dev)

```sh
# Install browsers for the first run
npx playwright install

# When testing on CI, must build the project first
npm run build

# Runs the end-to-end tests
npm run test:e2e
# Runs the tests only on Chromium
npm run test:e2e -- --project=chromium
# Runs the tests of a specific file
npm run test:e2e -- tests/example.spec.ts
# Runs the tests in debug mode
npm run test:e2e -- --debug
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
