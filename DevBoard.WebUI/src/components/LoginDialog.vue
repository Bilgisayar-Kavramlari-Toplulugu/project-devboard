<template>
  <Teleport to="body">
    <Transition name="dialog">
      <div v-if="modelValue" class="overlay" @click.self="close">
        <div class="dialog" role="dialog" aria-modal="true" aria-labelledby="dialog-title">

          <button class="close-btn" @click="close" aria-label="Close">
            <svg width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path d="M18 6 6 18M6 6l12 12"/>
            </svg>
          </button>

          <div class="dialog-logo">
            <span class="logo-icon">⬡</span>
            <span class="logo-text">DevBoard</span>
          </div>

          <h2 id="dialog-title" class="dialog-title">Welcome back</h2>
          <p class="dialog-sub">Log in to your account</p>

          <form class="form" @submit.prevent="handleSubmit">
            <!-- Email -->
            <div class="field">
              <label class="label">Email</label>
              <div class="input-wrap">
                <svg class="input-icon" width="15" height="15" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                  <rect x="2" y="4" width="20" height="16" rx="2"/>
                  <path d="m2 7 10 7 10-7"/>
                </svg>
                <input
                  v-model="form.email"
                  type="email"
                  class="input"
                  :class="{ error: errors.email }"
                  placeholder="you@example.com"
                  autocomplete="email"
                />
              </div>
              <span v-if="errors.email" class="field-error">{{ errors.email }}</span>
            </div>

            <!-- Password -->
            <div class="field">
              <label class="label">Password</label>
              <div class="input-wrap">
                <svg class="input-icon" width="15" height="15" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                  <rect x="3" y="11" width="18" height="11" rx="2"/>
                  <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
                </svg>
                <input
                  v-model="form.password"
                  type="password"
                  class="input"
                  :class="{ error: errors.password }"
                  placeholder="••••••••"
                  autocomplete="current-password"
                />
              </div>
              <span v-if="errors.password" class="field-error">{{ errors.password }}</span>
            </div>

            <button type="submit" class="submit-btn" :disabled="loading">
              <span v-if="!loading">Log In</span>
              <span v-else class="spinner"></span>
            </button>
          </form>

          <p class="dialog-footer">
            Don't have an account?
            <a href="#" class="footer-link" @click.prevent="openSignUp">Create one</a>
          </p>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'

const props = defineProps({
  modelValue: { type: Boolean, default: false }
})

const emit = defineEmits(['update:modelValue', 'open-signup'])

const loading = ref(false)
const form = reactive({ email: '', password: '' })
const errors = reactive({ email: '', password: '' })

function close() {
  emit('update:modelValue', false)
}

function openSignUp() {
  close()
  emit('open-signup')
}

watch(() => props.modelValue, (val) => {
  if (!val) {
    form.email = ''
    form.password = ''
    errors.email = ''
    errors.password = ''
    loading.value = false
  }
  if (val) {
    document.addEventListener('keydown', onKeydown)
    document.body.style.overflow = 'hidden'
  } else {
    document.removeEventListener('keydown', onKeydown)
    document.body.style.overflow = ''
  }
})

function onKeydown(e) {
  if (e.key === 'Escape') close()
}

function validate() {
  errors.email = ''
  errors.password = ''
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

  return valid
}

async function handleSubmit() {
  if (!validate()) return
  loading.value = true
  // auth logic will go here
  await new Promise(r => setTimeout(r, 1200))
  loading.value = false
}
</script>

<style scoped>
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

.dialog {
  position: relative;
  width: 100%;
  max-width: 420px;
  background: rgba(18, 14, 35, 0.92);
  border: 1px solid rgba(255, 255, 255, 0.09);
  border-radius: 20px;
  padding: 2.25rem 2rem;
  box-shadow:
    0 0 0 1px rgba(168, 85, 247, 0.08),
    0 24px 64px rgba(0, 0, 0, 0.6),
    0 0 80px rgba(124, 58, 237, 0.12);
}

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

.close-btn:hover {
  background: rgba(255,255,255,0.09);
  color: rgba(226,232,240,0.85);
}

.dialog-logo {
  display: flex;
  align-items: center;
  gap: 0.45rem;
  margin-bottom: 1.5rem;
}

.logo-icon { font-size: 1.3rem; color: #a855f7; }
.logo-text  { font-size: 1rem; font-weight: 700; color: #fff; letter-spacing: -0.02em; }

.dialog-title {
  font-size: 1.55rem;
  font-weight: 800;
  color: #fff;
  letter-spacing: -0.03em;
  margin-bottom: 0.3rem;
}

.dialog-sub {
  color: rgba(226,232,240,0.4);
  font-size: 0.85rem;
  margin-bottom: 1.75rem;
}

.form { display: flex; flex-direction: column; gap: 1.1rem; }

.field { display: flex; flex-direction: column; gap: 0.38rem; }

.label {
  font-size: 0.78rem;
  font-weight: 600;
  color: rgba(226,232,240,0.65);
  letter-spacing: 0.02em;
}

.input-wrap { position: relative; display: flex; align-items: center; }

.input-icon {
  position: absolute;
  left: 0.8rem;
  color: rgba(226,232,240,0.25);
  pointer-events: none;
}

.input {
  width: 100%;
  background: rgba(255,255,255,0.04);
  border: 1px solid rgba(255,255,255,0.08);
  border-radius: 10px;
  color: #e2e8f0;
  font-size: 0.875rem;
  font-family: inherit;
  padding: 0.62rem 0.85rem 0.62rem 2.4rem;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.input::placeholder { color: rgba(226,232,240,0.18); }

.input:focus {
  border-color: rgba(168,85,247,0.5);
  box-shadow: 0 0 0 3px rgba(168,85,247,0.1);
}

.input.error { border-color: rgba(239,68,68,0.5); }
.input.error:focus { box-shadow: 0 0 0 3px rgba(239,68,68,0.1); }

.field-error { font-size: 0.73rem; color: #f87171; }

.submit-btn {
  margin-top: 0.4rem;
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
  box-shadow: 0 0 22px rgba(168,85,247,0.35);
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 44px;
}

.submit-btn:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 0 34px rgba(168,85,247,0.55);
}

.submit-btn:disabled { opacity: 0.7; cursor: not-allowed; }

.spinner {
  width: 17px;
  height: 17px;
  border: 2px solid rgba(255,255,255,0.3);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin { to { transform: rotate(360deg); } }

.dialog-footer {
  text-align: center;
  margin-top: 1.35rem;
  font-size: 0.8rem;
  color: rgba(226,232,240,0.38);
}

.footer-link {
  color: #a855f7;
  text-decoration: none;
  font-weight: 600;
  margin-left: 0.2rem;
  transition: color 0.2s;
}

.footer-link:hover { color: #c084fc; }

.dialog-enter-active,
.dialog-leave-active { transition: opacity 0.22s ease; }

.dialog-enter-active .dialog,
.dialog-leave-active .dialog { transition: opacity 0.22s ease, transform 0.22s ease; }

.dialog-enter-from,
.dialog-leave-to { opacity: 0; }

.dialog-enter-from .dialog,
.dialog-leave-to .dialog {
  opacity: 0;
  transform: scale(0.96) translateY(10px);
}
</style>
