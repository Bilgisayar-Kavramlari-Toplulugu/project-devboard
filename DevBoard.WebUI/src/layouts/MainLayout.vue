<template>
  <div class="app-wrapper">
    <!-- Transparent Header -->
    <header class="header">
      <div class="header-inner">
        <div class="logo">
          <span class="logo-icon">⬡</span>
          <span class="logo-text">DevBoard</span>
        </div>

        <nav class="nav">
          <a href="#" class="nav-link">Explore</a>
          <a href="#" class="nav-link">Talents</a>
          <a href="#" class="nav-link">Blog</a>
          <a href="#" class="nav-link">Pricing</a>
        </nav>

        <div class="header-actions">
          <button class="btn-ghost" @click="openAuth('login')">Log in</button>
          <button class="btn-primary" @click="openAuth('signup')">Get Started</button>
        </div>
      </div>
    </header>

    <!-- Page Content -->
    <main class="content">
      <router-view />
    </main>

    <!-- Auth Dialog -->
    <AuthDialog v-model="authOpen" :initial-mode="authMode" />

    <!-- Footer -->
    <footer class="footer">
      <div class="footer-inner">
        <div class="footer-brand">
          <span class="logo-icon">⬡</span>
          <span class="logo-text">DevBoard</span>
          <p class="footer-desc">The developer showcase platform built for engineers, by engineers.</p>
        </div>

        <div class="footer-links">
          <div class="footer-col">
            <h4>Product</h4>
            <a href="#">Explore</a>
            <a href="#">Talents</a>
            <a href="#">Pricing</a>
          </div>
          <div class="footer-col">
            <h4>Company</h4>
            <a href="#">About</a>
            <a href="#">Blog</a>
            <a href="#">Careers</a>
          </div>
          <div class="footer-col">
            <h4>Legal</h4>
            <a href="#">Privacy</a>
            <a href="#">Terms</a>
            <a href="#">Cookies</a>
          </div>
        </div>
      </div>

      <div class="footer-bottom">
        <p>© 2026 DevBoard. All rights reserved.</p>
        <div class="socials">
          <a href="#">GitHub</a>
          <a href="#">Twitter</a>
          <a href="#">LinkedIn</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import AuthDialog from '../components/AuthDialog.vue'

const authOpen = ref(false)
const authMode = ref('login')

function openAuth(mode) {
  authMode.value = mode
  authOpen.value = true
}
</script>

<style scoped>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

.app-wrapper {
  min-height: 100vh;
  background: #0d0d1a;
  background-image:
    radial-gradient(ellipse 80% 60% at 50% -10%, rgba(120, 40, 200, 0.45) 0%, transparent 70%),
    radial-gradient(ellipse 40% 40% at 80% 20%, rgba(80, 20, 180, 0.3) 0%, transparent 60%);
  font-family: 'Inter', 'Segoe UI', sans-serif;
  color: #e2e8f0;
  display: flex;
  flex-direction: column;
}

/* ── HEADER ── */
.header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
  background: rgba(13, 13, 26, 0.45);
  backdrop-filter: blur(18px);
  -webkit-backdrop-filter: blur(18px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
}

.header-inner {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 2rem;
  height: 68px;
  display: flex;
  align-items: center;
  gap: 2rem;
}

.logo {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  text-decoration: none;
  flex-shrink: 0;
}

.logo-icon {
  font-size: 1.5rem;
  color: #a855f7;
}

.logo-text {
  font-size: 1.2rem;
  font-weight: 700;
  color: #fff;
  letter-spacing: -0.02em;
}

.nav {
  display: flex;
  gap: 0.25rem;
  flex: 1;
  justify-content: center;
}

.nav-link {
  color: rgba(226, 232, 240, 0.7);
  text-decoration: none;
  font-size: 0.9rem;
  padding: 0.4rem 0.9rem;
  border-radius: 8px;
  transition: all 0.2s;
}

.nav-link:hover {
  color: #fff;
  background: rgba(255, 255, 255, 0.07);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  flex-shrink: 0;
}

.btn-ghost {
  background: none;
  border: none;
  color: rgba(226, 232, 240, 0.8);
  font-size: 0.875rem;
  padding: 0.4rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-ghost:hover {
  color: #fff;
  background: rgba(255, 255, 255, 0.07);
}

.btn-primary {
  background: linear-gradient(135deg, #7c3aed, #a855f7);
  border: none;
  color: #fff;
  font-size: 0.875rem;
  font-weight: 600;
  padding: 0.5rem 1.2rem;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
  box-shadow: 0 0 20px rgba(168, 85, 247, 0.35);
}

.btn-primary:hover {
  transform: translateY(-1px);
  box-shadow: 0 0 30px rgba(168, 85, 247, 0.55);
}

/* ── CONTENT ── */
.content {
  flex: 1;
  padding-top: 68px;
}

/* ── FOOTER ── */
.footer {
  background: rgba(13, 13, 26, 0.8);
  border-top: 1px solid rgba(255, 255, 255, 0.06);
  padding: 3rem 2rem 1.5rem;
}

.footer-inner {
  max-width: 1280px;
  margin: 0 auto;
  display: flex;
  gap: 4rem;
  padding-bottom: 2.5rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
}

.footer-brand {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.footer-brand .logo-icon {
  font-size: 1.4rem;
  color: #a855f7;
}

.footer-desc {
  color: rgba(226, 232, 240, 0.45);
  font-size: 0.875rem;
  line-height: 1.6;
  max-width: 260px;
}

.footer-links {
  display: flex;
  gap: 4rem;
}

.footer-col {
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
}

.footer-col h4 {
  color: #fff;
  font-size: 0.85rem;
  font-weight: 600;
  margin-bottom: 0.25rem;
}

.footer-col a {
  color: rgba(226, 232, 240, 0.5);
  text-decoration: none;
  font-size: 0.85rem;
  transition: color 0.2s;
}

.footer-col a:hover {
  color: #a855f7;
}

.footer-bottom {
  max-width: 1280px;
  margin: 1.5rem auto 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.footer-bottom p {
  color: rgba(226, 232, 240, 0.35);
  font-size: 0.8rem;
}

.socials {
  display: flex;
  gap: 1.25rem;
}

.socials a {
  color: rgba(226, 232, 240, 0.4);
  text-decoration: none;
  font-size: 0.8rem;
  transition: color 0.2s;
}

.socials a:hover {
  color: #a855f7;
}
</style>
