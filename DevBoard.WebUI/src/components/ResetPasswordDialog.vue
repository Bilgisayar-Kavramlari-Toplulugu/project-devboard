<template>
  <Teleport to="body">
    <Transition name="fade">
      <div v-if="modelValue" class="overlay" @click.self="close">
        <div class="dialog" role="dialog" aria-modal="true" aria-labelledby="reset-dialog-title">

          <button class="close-btn" @click="close" aria-label="Close">
            <svg width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path d="M18 6 6 18M6 6l12 12"/>
            </svg>
          </button>

          <!-- Invalid / Missing Token State -->
          <template v-if="!token">
            <div class="icon-wrap error-color">
              <svg width="48" height="48" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10"/>
                <path d="M12 8v4M12 16h.01"/>
              </svg>
            </div>
            <h2 id="reset-dialog-title" class="dialog-title">Invalid link</h2>
            <p class="dialog-sub">This password reset link is missing a token. Please request a new one.</p>
            <button class="submit-btn" style="margin-top: 0.5rem;" @click="openForgot">
              Request New Link
            </button>
          </template>

          <!-- Success State -->
          <template v-else-if="success">
            <div class="icon-wrap success-color">
              <svg width="48" height="48" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10"/>
                <path d="m9 12 2 2 4-4"/>
              </svg>
            </div>
            <h2 id="reset-dialog-title" class="dialog-title">Password updated!</h2>
            <p class="dialog-sub">
              Your password has been reset successfully. Redirecting you to sign in
              <template v-if="countdown > 0"> in {{ countdown }}...</template>
            </p>
            <button class="submit-btn" style="margin-top: 0.5rem;" @click="openLogin">
              Sign In Now
            </button>
          </template>

          <!-- Form State -->
          <template v-else>
            <h2 id="reset-dialog-title" class="dialog-title">Reset password</h2>
            <p class="dialog-sub">Enter and confirm your new password below.</p>

            <!-- Global error banner -->
            <div v-if="globalError" class="error-banner">
              <svg width="16" height="16" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10"/>
                <path d="M12 8v4M12 16h.01"/>
              </svg>
              <span>{{ globalError }}</span>
            </div>

            <form class="form" @submit.prevent="handleSubmit">
              <!-- New Password -->
              <div class="field">
                <label class="label" for="reset-password">New Password</label>
                <div class="input-wrap">
                  <svg class="input-icon" width="15" height="15" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                    <rect x="3" y="11" width="18" height="11" rx="2"/>
                    <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
                  </svg>
                  <input
                    id="reset-password"
                    v-model="form.password"
                    :type="showPassword ? 'text' : 'password'"
                    class="input input-with-toggle"
                    placeholder="••••••••"
                    autocomplete="new-password"
                    :class="{ error: errors.password }"
                    @input="errors.password = ''"
                  />
                  <button type="button" class="eye-btn" @click="showPassword = !showPassword" tabindex="-1">
                    <svg v-if="showPassword" width="17" height="17" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                      <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/>
                    </svg>
                    <svg v-else width="17" height="17" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                      <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94"/>
                      <path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19"/>
                      <line x1="1" y1="1" x2="23" y2="23"/>
                    </svg>
                  </button>
                </div>
                <span v-if="errors.password" class="field-error">{{ errors.password }}</span>
              </div>

              <!-- Confirm Password -->
              <div class="field">
                <label class="label" for="reset-confirm-password">Confirm Password</label>
                <div class="input-wrap">
                  <svg class="input-icon" width="15" height="15" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                    <rect x="3" y="11" width="18" height="11" rx="2"/>
                    <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
                  </svg>
                  <input
                    id="reset-confirm-password"
                    v-model="form.confirmPassword"
                    :type="showConfirm ? 'text' : 'password'"
                    class="input input-with-toggle"
                    placeholder="••••••••"
                    autocomplete="new-password"
                    :class="{ error: errors.confirmPassword }"
                    @input="errors.confirmPassword = ''"
                  />
                  <button type="button" class="eye-btn" @click="showConfirm = !showConfirm" tabindex="-1">
                    <svg v-if="showConfirm" width="17" height="17" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                      <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/>
                    </svg>
                    <svg v-else width="17" height="17" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
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
                <span v-if="!loading">Reset Password</span>
                <span v-else class="spinner"></span>
              </button>
            </form>

            <p class="dialog-footer">
              Remembered your password?
              <button class="switch-link" @click="openLogin">Sign in</button>
            </p>
          </template>

        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { ref, reactive, watch, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const props = defineProps({
  modelValue: { type: Boolean, default: false }
})

const emit = defineEmits(['update:modelValue', 'open-login', 'open-forgot'])

const route = useRoute()
const router = useRouter()

const token = ref('')
const showPassword = ref(false)
const showConfirm = ref(false)
const loading = ref(false)
const success = ref(false)
const countdown = ref(3)
const globalError = ref('')

const form = reactive({ password: '', confirmPassword: '' })
const errors = reactive({ password: '', confirmPassword: '' })

// AbortController — in-flight isteği iptal etmek için
let abortController = null

// Countdown timer referansı — dialog kapandığında temizlemek için
let redirectTimer = null

const API_BASE = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080/api/v1'

function close() {
  emit('update:modelValue', false)
  if (route.name === 'ResetPassword') {
    router.push('/')
  }
}

function openLogin() {
  emit('update:modelValue', false)
  emit('open-login')
  if (route.name === 'ResetPassword') {
    router.push('/')
  }
}

function openForgot() {
  emit('update:modelValue', false)
  emit('open-forgot')
  if (route.name === 'ResetPassword') {
    router.push('/')
  }
}

function onKeydown(e) {
  if (e.key === 'Escape') close()
}

onUnmounted(() => document.removeEventListener('keydown', onKeydown))

watch(() => props.modelValue, (val) => {
  if (val) {
    token.value = route.query.token ?? ''
    document.addEventListener('keydown', onKeydown)
    document.body.style.overflow = 'hidden'
  } else {
    document.removeEventListener('keydown', onKeydown)
    document.body.style.overflow = ''
    // Dialog kapandığında devam eden fetch ve timer'ı temizle
    abortController?.abort()
    abortController = null
    if (redirectTimer) {
      clearInterval(redirectTimer)
      redirectTimer = null
    }
    form.password = ''
    form.confirmPassword = ''
    errors.password = ''
    errors.confirmPassword = ''
    globalError.value = ''
    success.value = false
    countdown.value = 3
    showPassword.value = false
    showConfirm.value = false
    loading.value = false
  }
}, { immediate: true })

function validate() {
  errors.password = ''
  errors.confirmPassword = ''
  let valid = true

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

function startRedirectCountdown() {
  // Önceki timer varsa temizle
  if (redirectTimer) {
    clearInterval(redirectTimer)
  }
  redirectTimer = setInterval(() => {
    countdown.value -= 1
    if (countdown.value <= 0) {
      clearInterval(redirectTimer)
      redirectTimer = null
      openLogin()
    }
  }, 1000)
}

async function handleSubmit() {
  if (!validate()) return

  // Önceki isteği iptal et, yeni controller oluştur
  abortController?.abort()
  abortController = new AbortController()

  loading.value = true
  globalError.value = ''

  try {
    const res = await fetch(`${API_BASE}/auth/reset-password`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        token: token.value,
        new_password: form.password,
      }),
      signal: abortController.signal,
    })

    if (res.ok) {
      success.value = true
      startRedirectCountdown()
    } else {
      const data = await res.json().catch(() => ({}))
      globalError.value = data?.error?.message ?? 'Something went wrong. The link may have expired.'
    }
  } catch (err) {
    // Dialog kapandığında abort edilen isteği sessizce geç
    if (err.name === 'AbortError') return
    globalError.value = 'Unable to reach the server. Please check your connection.'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* ── OVERLAY ── */
.overlay {
  position: fixed;
  inset: 0;
  z-index: 200;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1.5rem;
  background: rgba(4, 4, 15, 0.65);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
}

/* ── DIALOG ── */
.dialog {
  position: relative;
  width: 100%;
  max-width: 420px;
  background: rgba(18, 14, 35, 0.95);
  border: 1px solid rgba(255, 255, 255, 0.09);
  border-radius: 20px;
  padding: 2.25rem 2rem;
  box-shadow:
    0 0 0 1px rgba(168, 85, 247, 0.08),
    0 24px 64px rgba(0, 0, 0, 0.6),
    0 0 80px rgba(124, 58, 237, 0.12);
}

/* ── CLOSE ── */
.close-btn {
  position: absolute;
  top: 1rem;
  right: 1rem;
  background: rgba(255,255,255,0.05);
  border: 1px solid rgba(255,255,255,0.08);
  color: rgba(226,232,240,0.4);
  border-radius: 8px;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
}
.close-btn:hover { background: rgba(255,255,255,0.09); color: rgba(226,232,240,0.85); }

/* ── ICONS ── */
.icon-wrap {
  display: flex;
  justify-content: center;
  margin-bottom: 1.25rem;
}
.success-color { color: #4ade80; }
.error-color   { color: #f87171; }

/* ── TITLES ── */
.dialog-title {
  font-size: 1.45rem;
  font-weight: 800;
  color: #fff;
  letter-spacing: -0.03em;
  margin-bottom: 0.3rem;
}

.dialog-sub {
  color: rgba(226, 232, 240, 0.4);
  font-size: 0.85rem;
  margin-bottom: 1.75rem;
  line-height: 1.6;
}

/* ── ERROR BANNER ── */
.error-banner {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: rgba(239, 68, 68, 0.08);
  border: 1px solid rgba(239, 68, 68, 0.25);
  border-radius: 10px;
  padding: 0.65rem 0.875rem;
  color: #f87171;
  font-size: 0.83rem;
  margin-bottom: 1.25rem;
}

/* ── FORM ── */
.form { display: flex; flex-direction: column; gap: 1rem; }

.field { display: flex; flex-direction: column; gap: 0.38rem; }

.label {
  font-size: 0.78rem;
  font-weight: 600;
  color: rgba(226, 232, 240, 0.65);
  letter-spacing: 0.02em;
}

.input-wrap { position: relative; display: flex; align-items: center; }

.input-icon {
  position: absolute;
  left: 0.8rem;
  color: rgba(226, 232, 240, 0.25);
  pointer-events: none;
}

.input {
  width: 100%;
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 10px;
  color: #e2e8f0;
  font-size: 0.875rem;
  font-family: inherit;
  padding: 0.62rem 0.85rem 0.62rem 2.4rem;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.input::placeholder { color: rgba(226, 232, 240, 0.18); }
.input:focus { border-color: rgba(168, 85, 247, 0.5); box-shadow: 0 0 0 3px rgba(168, 85, 247, 0.1); }
.input.error { border-color: rgba(239, 68, 68, 0.5); }
.input.error:focus { box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1); }
.input-with-toggle { padding-right: 2.6rem; }

.eye-btn {
  position: absolute;
  right: 0.7rem;
  background: none;
  border: none;
  color: rgba(226, 232, 240, 0.3);
  cursor: pointer;
  padding: 0.2rem;
  display: flex;
  align-items: center;
  border-radius: 5px;
  transition: color 0.2s;
}
.eye-btn:hover { color: rgba(226, 232, 240, 0.75); }

.field-error { font-size: 0.73rem; color: #f87171; }

/* ── SUBMIT ── */
.submit-btn {
  margin-top: 0.25rem;
  width: 100%;
  background: linear-gradient(135deg, #7c3aed, #a855f7);
  border: none;
  color: #fff;
  font-size: 0.9rem;
  font-weight: 600;
  font-family: inherit;
  padding: 0.7rem;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0 0 22px rgba(168, 85, 247, 0.35);
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 44px;
}
.submit-btn:hover:not(:disabled) { transform: translateY(-1px); box-shadow: 0 0 34px rgba(168, 85, 247, 0.55); }
.submit-btn:disabled { opacity: 0.7; cursor: not-allowed; }

.spinner {
  width: 17px; height: 17px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }

/* ── FOOTER ── */
.dialog-footer {
  text-align: center;
  margin-top: 1.25rem;
  font-size: 0.8rem;
  color: rgba(226, 232, 240, 0.38);
}

.switch-link {
  background: none;
  border: none;
  color: #a855f7;
  font-size: 0.8rem;
  font-weight: 600;
  font-family: inherit;
  margin-left: 0.2rem;
  cursor: pointer;
  padding: 0;
  transition: color 0.2s;
}
.switch-link:hover { color: #c084fc; }

/* ── FADE TRANSITION ── */
.fade-enter-active, .fade-leave-active { transition: opacity 0.22s ease; }
.fade-enter-active .dialog, .fade-leave-active .dialog { transition: opacity 0.22s ease, transform 0.22s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
.fade-enter-from .dialog, .fade-leave-to .dialog { opacity: 0; transform: scale(0.96) translateY(8px); }
</style>
