<template>
  <div class="login-page">
    <!-- Background glow blobs -->
    <div class="blob blob-1"></div>
    <div class="blob blob-2"></div>

    <div class="login-card">
      <!-- Logo -->
      <div class="card-logo">
        <span class="logo-icon">⬡</span>
        <span class="logo-text">DevBoard</span>
      </div>

      <h1 class="card-title">Welcome back</h1>
      <p class="card-sub">Sign in to your account</p>

      <form class="form" @submit.prevent="handleSubmit">
        <!-- Email -->
        <div class="field">
          <label class="label">Email</label>
          <div class="input-wrap">
            <svg class="input-icon" width="16" height="16" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
              <path d="M4 4h16v16H4z" stroke="none"/>
              <rect x="2" y="4" width="20" height="16" rx="2"/>
              <path d="m2 7 10 7 10-7"/>
            </svg>
            <input
              v-model="form.email"
              type="email"
              class="input"
              placeholder="you@example.com"
              autocomplete="email"
              :class="{ error: errors.email }"
            />
          </div>
          <span v-if="errors.email" class="field-error">{{ errors.email }}</span>
        </div>

        <!-- Password -->
        <div class="field">
          <label class="label">Password</label>
          <div class="input-wrap">
            <svg class="input-icon" width="16" height="16" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
              <rect x="3" y="11" width="18" height="11" rx="2"/>
              <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
            </svg>
            <input
              v-model="form.password"
              type="password"
              class="input"
              placeholder="••••••••"
              autocomplete="current-password"
              :class="{ error: errors.password }"
            />
          </div>
          <span v-if="errors.password" class="field-error">{{ errors.password }}</span>
        </div>

        <!-- Confirm Password -->
        <div class="field">
          <label class="label">Confirm Password</label>
          <div class="input-wrap">
            <svg class="input-icon" width="16" height="16" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
              <rect x="3" y="11" width="18" height="11" rx="2"/>
              <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
            </svg>
            <input
              v-model="form.confirmPassword"
              :type="showConfirm ? 'text' : 'password'"
              class="input input-with-toggle"
              placeholder="••••••••"
              autocomplete="new-password"
              :class="{ error: errors.confirmPassword }"
            />
            <button type="button" class="eye-btn" @click="showConfirm = !showConfirm" tabindex="-1">
              <!-- Eye open -->
              <svg v-if="showConfirm" width="18" height="18" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                <circle cx="12" cy="12" r="3"/>
              </svg>
              <!-- Eye closed -->
              <svg v-else width="18" height="18" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94"/>
                <path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19"/>
                <line x1="1" y1="1" x2="23" y2="23"/>
              </svg>
            </button>
          </div>
          <span v-if="errors.confirmPassword" class="field-error">{{ errors.confirmPassword }}</span>
        </div>

        <!-- Submit -->
        <button type="submit" class="submit-btn" :disabled="loading">
          <span v-if="!loading">Sign In</span>
          <span v-else class="spinner"></span>
        </button>
      </form>

      <p class="card-footer">
        Don't have an account?
        <a href="#" class="footer-link">Create one</a>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const showConfirm = ref(false)
const loading = ref(false)

const form = reactive({
  email: '',
  password: '',
  confirmPassword: '',
})

const errors = reactive({
  email: '',
  password: '',
  confirmPassword: '',
})

function validate() {
  errors.email = ''
  errors.password = ''
  errors.confirmPassword = ''

  let valid = true

  if (!form.email) {
    errors.email = 'Email is required.'
    valid = false
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email)) {
    errors.email = 'Please enter a valid email.'
    valid = false
  }

  if (!form.password) {
    errors.password = 'Password is required.'
    valid = false
  } else if (form.password.length < 6) {
    errors.password = 'Password must be at least 6 characters.'
    valid = false
  }

  if (!form.confirmPassword) {
    errors.confirmPassword = 'Please confirm your password.'
    valid = false
  } else if (form.password !== form.confirmPassword) {
    errors.confirmPassword = 'Passwords do not match.'
    valid = false
  }

  return valid
}

async function handleSubmit() {
  if (!validate()) return

  loading.value = true
  // Placeholder — auth logic will go here
  await new Promise(r => setTimeout(r, 1200))
  loading.value = false
}
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  background: #0d0d1a;
  background-image:
    radial-gradient(ellipse 80% 60% at 50% -10%, rgba(120, 40, 200, 0.45) 0%, transparent 70%),
    radial-gradient(ellipse 40% 40% at 80% 20%, rgba(80, 20, 180, 0.3) 0%, transparent 60%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  font-family: 'Inter', 'Segoe UI', sans-serif;
  position: relative;
  overflow: hidden;
}

/* decorative blobs */
.blob {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  pointer-events: none;
}

.blob-1 {
  width: 400px;
  height: 400px;
  background: rgba(124, 58, 237, 0.18);
  top: -100px;
  left: -100px;
}

.blob-2 {
  width: 350px;
  height: 350px;
  background: rgba(236, 72, 153, 0.1);
  bottom: -80px;
  right: -80px;
}

/* ── CARD ── */
.login-card {
  position: relative;
  width: 100%;
  max-width: 420px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 20px;
  padding: 2.5rem 2.25rem;
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
}

.card-logo {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 2rem;
}

.logo-icon {
  font-size: 1.4rem;
  color: #a855f7;
}

.logo-text {
  font-size: 1.1rem;
  font-weight: 700;
  color: #fff;
  letter-spacing: -0.02em;
}

.card-title {
  font-size: 1.7rem;
  font-weight: 800;
  color: #fff;
  letter-spacing: -0.03em;
  margin-bottom: 0.35rem;
}

.card-sub {
  color: rgba(226, 232, 240, 0.45);
  font-size: 0.875rem;
  margin-bottom: 2rem;
}

/* ── FORM ── */
.form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.label {
  font-size: 0.8rem;
  font-weight: 600;
  color: rgba(226, 232, 240, 0.7);
  letter-spacing: 0.02em;
}

.input-wrap {
  position: relative;
  display: flex;
  align-items: center;
}

.input-icon {
  position: absolute;
  left: 0.875rem;
  color: rgba(226, 232, 240, 0.3);
  pointer-events: none;
  flex-shrink: 0;
}

.input {
  width: 100%;
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 10px;
  color: #e2e8f0;
  font-size: 0.9rem;
  font-family: inherit;
  padding: 0.65rem 0.9rem 0.65rem 2.6rem;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.input::placeholder {
  color: rgba(226, 232, 240, 0.2);
}

.input:focus {
  border-color: rgba(168, 85, 247, 0.5);
  box-shadow: 0 0 0 3px rgba(168, 85, 247, 0.1);
}

.input.error {
  border-color: rgba(239, 68, 68, 0.5);
}

.input.error:focus {
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1);
}

.input-with-toggle {
  padding-right: 2.75rem;
}

.eye-btn {
  position: absolute;
  right: 0.75rem;
  background: none;
  border: none;
  color: rgba(226, 232, 240, 0.35);
  cursor: pointer;
  padding: 0.25rem;
  display: flex;
  align-items: center;
  border-radius: 6px;
  transition: color 0.2s;
}

.eye-btn:hover {
  color: rgba(226, 232, 240, 0.75);
}

.field-error {
  font-size: 0.75rem;
  color: #f87171;
}

/* ── SUBMIT ── */
.submit-btn {
  margin-top: 0.5rem;
  width: 100%;
  background: linear-gradient(135deg, #7c3aed, #a855f7);
  border: none;
  color: #fff;
  font-size: 0.95rem;
  font-weight: 600;
  font-family: inherit;
  padding: 0.75rem;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0 0 24px rgba(168, 85, 247, 0.35);
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 46px;
}

.submit-btn:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 0 36px rgba(168, 85, 247, 0.55);
}

.submit-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

/* loading spinner */
.spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255,255,255,0.3);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* ── FOOTER ── */
.card-footer {
  text-align: center;
  margin-top: 1.5rem;
  font-size: 0.83rem;
  color: rgba(226, 232, 240, 0.4);
}

.footer-link {
  color: #a855f7;
  text-decoration: none;
  font-weight: 600;
  margin-left: 0.25rem;
  transition: color 0.2s;
}

.footer-link:hover {
  color: #c084fc;
}
</style>
