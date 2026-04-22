<template>
  <Teleport to="body">
    <Transition name="fade">
      <div v-if="modelValue" class="overlay" @click.self="close">
        <div class="dialog" role="dialog" aria-modal="true" aria-labelledby="forgot-dialog-title">

          <button class="close-btn" @click="close" aria-label="Close">
            <svg width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path d="M18 6 6 18M6 6l12 12"/>
            </svg>
          </button>

          <!-- Success State -->
          <template v-if="success">
            <div class="icon-wrap success-color">
              <svg width="48" height="48" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10"/>
                <path d="m9 12 2 2 4-4"/>
              </svg>
            </div>
            <h2 id="forgot-dialog-title" class="dialog-title">Check your email</h2>
            <p class="dialog-sub">
              If an account exists for <strong>{{ submittedEmail }}</strong>, we've sent a password reset link. Check your inbox and spam folder.
            </p>
            <button class="submit-btn" style="margin-top: 1.5rem;" @click="openLogin">
              Back to Sign In
            </button>
          </template>

          <!-- Form State -->
          <template v-else>
            <h2 id="forgot-dialog-title" class="dialog-title">Forgot password?</h2>
            <p class="dialog-sub">Enter your email and we'll send you a reset link.</p>

            <!-- Global error banner -->
            <div v-if="globalError" class="error-banner">
              <svg width="16" height="16" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10"/>
                <path d="M12 8v4M12 16h.01"/>
              </svg>
              <span>{{ globalError }}</span>
            </div>

            <form class="form" @submit.prevent="handleSubmit">
              <!-- Email -->
              <div class="field">
                <label class="label" for="forgot-email">Email</label>
                <div class="input-wrap">
                  <svg class="input-icon" width="15" height="15" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                    <rect x="2" y="4" width="20" height="16" rx="2"/>
                    <path d="m2 7 10 7 10-7"/>
                  </svg>
                  <input
                    id="forgot-email"
                    v-model="email"
                    type="email"
                    class="input"
                    placeholder="you@example.com"
                    autocomplete="email"
                    :class="{ error: emailError }"
                    @input="emailError = ''"
                  />
                </div>
                <span v-if="emailError" class="field-error">{{ emailError }}</span>
              </div>

              <!-- Submit -->
              <button type="submit" class="submit-btn" :disabled="loading">
                <span v-if="!loading">Send Reset Link</span>
                <span v-else class="spinner"></span>
              </button>
            </form>

            <p class="dialog-footer">
              Remember your password?
              <button class="switch-link" @click="openLogin">Sign in</button>
            </p>
          </template>

        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: { type: Boolean, default: false }
})

const emit = defineEmits(['update:modelValue', 'open-login'])

const email = ref('')
const emailError = ref('')
const globalError = ref('')
const loading = ref(false)
const success = ref(false)
const submittedEmail = ref('')

// AbortController — in-flight isteği iptal etmek için
let abortController = null

const API_BASE = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080/api/v1'

function close() {
  emit('update:modelValue', false)
}

function openLogin() {
  close()
  emit('open-login')
}

watch(() => props.modelValue, (val) => {
  if (val) {
    document.addEventListener('keydown', onKeydown)
    document.body.style.overflow = 'hidden'
  } else {
    document.removeEventListener('keydown', onKeydown)
    document.body.style.overflow = ''
    // Dialog kapandığında devam eden fetch varsa iptal et
    abortController?.abort()
    abortController = null
    email.value = ''
    emailError.value = ''
    globalError.value = ''
    success.value = false
    submittedEmail.value = ''
    loading.value = false
  }
})

function onKeydown(e) {
  if (e.key === 'Escape') close()
}

function validate() {
  emailError.value = ''
  if (!email.value) {
    emailError.value = 'Email is required.'
    return false
  }
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) {
    emailError.value = 'Please enter a valid email.'
    return false
  }
  return true
}

async function handleSubmit() {
  if (!validate()) return

  // Önceki isteği iptal et, yeni controller oluştur
  abortController?.abort()
  abortController = new AbortController()

  loading.value = true
  globalError.value = ''

  try {
    const res = await fetch(`${API_BASE}/auth/forgot-password`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email: email.value }),
      signal: abortController.signal,
    })

    if (res.ok) {
      submittedEmail.value = email.value
      success.value = true
    } else {
      const data = await res.json().catch(() => ({}))
      globalError.value = data?.error?.message ?? 'Something went wrong. Please try again.'
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

/* ── SUCCESS ICON ── */
.icon-wrap {
  display: flex;
  justify-content: center;
  margin-bottom: 1.25rem;
}
.success-color { color: #4ade80; }

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

.dialog-sub strong {
  color: rgba(226, 232, 240, 0.75);
  font-weight: 600;
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
