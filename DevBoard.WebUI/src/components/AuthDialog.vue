<template>
  <Teleport to="body">
    <Transition name="fade">
      <div v-if="modelValue" class="overlay" @click.self="close">
        <div class="dialog" role="dialog" aria-modal="true">

          <button class="close-btn" @click="close" aria-label="Close">
            <svg width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path d="M18 6 6 18M6 6l12 12"/>
            </svg>
          </button>

          <div class="track-wrap">
            <div class="track" :class="{ 'show-signup': mode === 'signup' }">

              <!-- LOG IN -->
              <div class="panel" aria-hidden="mode !== 'login'">
                <div class="panel-head">
                  <h2 class="dialog-title">Welcome back</h2>
                  <p class="dialog-sub">Log in to your account</p>
                </div>

                <form class="form" @submit.prevent="handleLogin">
                  <div v-if="loginApiError" class="api-message api-error">
                    {{ loginApiError }}
                  </div>

                  <div class="field">
                    <label class="label">Email</label>
                    <div class="input-wrap">
                      <EmailIcon />
                      <input
                        v-model="login.email"
                        type="email"
                        class="input"
                        :class="{ error: loginErr.email }"
                        placeholder="you@example.com"
                        autocomplete="email"
                        :tabindex="mode === 'login' ? 0 : -1"
                      />
                    </div>
                    <span v-if="loginErr.email" class="field-error">{{ loginErr.email }}</span>
                  </div>

                  <div class="field">
                    <label class="label">Password</label>
                    <div class="input-wrap">
                      <LockIcon />
                      <input
                        v-model="login.password"
                        type="password"
                        class="input"
                        :class="{ error: loginErr.password }"
                        placeholder="••••••••"
                        autocomplete="current-password"
                        :tabindex="mode === 'login' ? 0 : -1"
                      />
                    </div>
                    <span v-if="loginErr.password" class="field-error">{{ loginErr.password }}</span>
                  </div>

                  <div class="field field-spacer" aria-hidden="true">
                    <label class="label">&nbsp;</label>
                    <div class="input-wrap">
                      <input class="input" tabindex="-1" disabled placeholder=" " />
                    </div>
                  </div>

                  <button type="submit" class="submit-btn" :disabled="loginLoading" :tabindex="mode === 'login' ? 0 : -1">
                    <span v-if="!loginLoading">Log In</span>
                    <span v-else class="spinner"></span>
                  </button>
                </form>

                <p class="switch-hint">
                  Don't have an account?
                  <button class="switch-link" @click="switchTo('signup')" :tabindex="mode === 'login' ? 0 : -1">Create one</button>
                </p>
              </div>

              <!-- SIGN UP -->
              <div class="panel" aria-hidden="mode !== 'signup'">
                <div class="panel-head">
                  <h2 class="dialog-title">Create account</h2>
                  <p class="dialog-sub">Join 20k+ developers on DevBoard</p>
                </div>

                <form class="form" @submit.prevent="handleSignup">
                  <div v-if="signupApiError" class="api-message api-error">
                    {{ signupApiError }}
                  </div>

                  <div v-if="signupApiSuccess" class="api-message api-success">
                    {{ signupApiSuccess }}
                  </div>

                  <div class="name-row">
                    <div class="field">
                      <label class="label">First Name</label>
                      <div class="input-wrap">
                        <UserIcon />
                        <input
                          v-model="signup.firstName"
                          type="text"
                          class="input"
                          :class="{ error: signupErr.firstName }"
                          placeholder="John"
                          autocomplete="given-name"
                          :tabindex="mode === 'signup' ? 0 : -1"
                        />
                      </div>
                      <span v-if="signupErr.firstName" class="field-error">{{ signupErr.firstName }}</span>
                    </div>

                    <div class="field">
                      <label class="label">Last Name</label>
                      <div class="input-wrap">
                        <UserIcon />
                        <input
                          v-model="signup.lastName"
                          type="text"
                          class="input"
                          :class="{ error: signupErr.lastName }"
                          placeholder="Doe"
                          autocomplete="family-name"
                          :tabindex="mode === 'signup' ? 0 : -1"
                        />
                      </div>
                      <span v-if="signupErr.lastName" class="field-error">{{ signupErr.lastName }}</span>
                    </div>
                  </div>

                  <div class="field">
                    <label class="label">Email</label>
                    <div class="input-wrap">
                      <EmailIcon />
                      <input
                        v-model="signup.email"
                        type="email"
                        class="input"
                        :class="{ error: signupErr.email }"
                        placeholder="you@example.com"
                        autocomplete="email"
                        :tabindex="mode === 'signup' ? 0 : -1"
                      />
                    </div>
                    <span v-if="signupErr.email" class="field-error">{{ signupErr.email }}</span>
                  </div>

                  <div class="field">
                    <label class="label">Role</label>
                    <div class="input-wrap">
                      <UserIcon />
                      <select
                        v-model="signup.role"
                        class="input select-input"
                        :class="{ error: signupErr.role }"
                        :disabled="rolesLoading"
                        :tabindex="mode === 'signup' ? 0 : -1"
                      >
                        <option value="" disabled>
                          {{ rolesLoading ? 'Loading roles...' : 'Select role' }}
                        </option>
                        <option v-for="role in roles" :key="role.id" :value="role.name">
                          {{ role.name }}
                        </option>
                      </select>
                    </div>
                    <span v-if="signupErr.role" class="field-error">{{ signupErr.role }}</span>
                  </div>

                  <div class="field">
                    <label class="label">Password</label>
                    <div class="input-wrap">
                      <LockIcon />
                      <input
                        v-model="signup.password"
                        type="password"
                        class="input"
                        :class="{ error: signupErr.password }"
                        placeholder="••••••••"
                        autocomplete="new-password"
                        :tabindex="mode === 'signup' ? 0 : -1"
                      />
                    </div>
                    <span v-if="signupErr.password" class="field-error">{{ signupErr.password }}</span>
                  </div>

                  <div class="field">
                    <label class="label">Confirm Password</label>
                    <div class="input-wrap">
                      <LockIcon />
                      <input
                        v-model="signup.confirmPassword"
                        :type="showConfirm ? 'text' : 'password'"
                        class="input input-with-toggle"
                        :class="{ error: signupErr.confirmPassword }"
                        placeholder="••••••••"
                        autocomplete="new-password"
                        :tabindex="mode === 'signup' ? 0 : -1"
                      />
                      <button type="button" class="eye-btn" @click="showConfirm = !showConfirm" tabindex="-1">
                        <svg v-if="showConfirm" width="17" height="17" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                          <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                          <circle cx="12" cy="12" r="3"/>
                        </svg>
                        <svg v-else width="17" height="17" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                          <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94"/>
                          <path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19"/>
                          <line x1="1" y1="1" x2="23" y2="23"/>
                        </svg>
                      </button>
                    </div>
                    <span v-if="signupErr.confirmPassword" class="field-error">{{ signupErr.confirmPassword }}</span>
                  </div>

                  <button type="submit" class="submit-btn" :disabled="signupLoading || rolesLoading" :tabindex="mode === 'signup' ? 0 : -1">
                    <span v-if="!signupLoading">Create Account</span>
                    <span v-else class="spinner"></span>
                  </button>
                </form>

                <p class="switch-hint">
                  Already have an account?
                  <button class="switch-link" @click="switchTo('login')" :tabindex="mode === 'signup' ? 0 : -1">Log in</button>
                </p>
              </div>

            </div>
          </div>

        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { ref, reactive, watch, onBeforeUnmount } from 'vue'
import { fetchRoles, loginUser, signupUser } from '../services/authService'

const EmailIcon = {
  template: `<svg class="input-icon" width="15" height="15" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24"><rect x="2" y="4" width="20" height="16" rx="2"/><path d="m2 7 10 7 10-7"/></svg>`
}
const LockIcon = {
  template: `<svg class="input-icon" width="15" height="15" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24"><rect x="3" y="11" width="18" height="11" rx="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>`
}

const UserIcon = {
  template: `<svg class="input-icon" width="15" height="15" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24"><path d="M12 12a4 4 0 1 0 0-8 4 4 0 0 0 0 8Z"/><path d="M4 20a8 8 0 1 1 16 0"/></svg>`
}

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  initialMode: { type: String, default: 'login' }
})
const emit = defineEmits(['update:modelValue'])

const mode = ref(props.initialMode)
const showConfirm = ref(false)

const login = reactive({ email: '', password: '' })
const loginErr = reactive({ email: '', password: '' })
const loginLoading = ref(false)
const loginApiError = ref('')

const signup = reactive({
  firstName: '',
  lastName: '',
  email: '',
  role: '',
  password: '',
  confirmPassword: ''
})

const signupErr = reactive({
  firstName: '',
  lastName: '',
  email: '',
  role: '',
  password: '',
  confirmPassword: ''
})

const signupLoading = ref(false)
const signupApiError = ref('')
const signupApiSuccess = ref('')

const roles = ref([])
const rolesLoading = ref(false)

function switchTo(target) {
  mode.value = target
}

function close() {
  emit('update:modelValue', false)
}

function onKeydown(e) {
  if (e.key === 'Escape') close()
}

async function loadRoles() {
  rolesLoading.value = true
  try {
    roles.value = await fetchRoles()
  } catch (error) {
    signupApiError.value = error.message || 'Failed to load roles.'
  } finally {
    rolesLoading.value = false
  }
}

watch(() => props.modelValue, async (val) => {
  if (val) {
    mode.value = props.initialMode
    document.addEventListener('keydown', onKeydown)
    document.body.style.overflow = 'hidden'
    if (roles.value.length === 0) {
  await loadRoles()
}
  } else {
    document.removeEventListener('keydown', onKeydown)
    document.body.style.overflow = ''
    Object.assign(login, { email: '', password: '' })
    Object.assign(loginErr, { email: '', password: '' })
    Object.assign(signup, {
      firstName: '',
      lastName: '',
      email: '',
      role: '',
      password: '',
      confirmPassword: ''
    })
    Object.assign(signupErr, {
      firstName: '',
      lastName: '',
      email: '',
      role: '',
      password: '',
      confirmPassword: ''
    })
    loginApiError.value = ''
    signupApiError.value = ''
    signupApiSuccess.value = ''
    showConfirm.value = false
    loginLoading.value = false
    signupLoading.value = false
  }
})

onBeforeUnmount(() => {
  document.removeEventListener('keydown', onKeydown)
  document.body.style.overflow = ''
})

function validateEmail(email, err) {
  if (!email) {
    err.email = 'Email is required.'
    return false
  }
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
    err.email = 'Please enter a valid email.'
    return false
  }
  err.email = ''
  return true
}

function validatePassword(password, err) {
  if (!password) {
    err.password = 'Password is required.'
    return false
  }
  if (password.length < 8) {
    err.password = 'Password must be at least 8 characters.'
    return false
  }
  err.password = ''
  return true
}

async function handleLogin() {
  loginApiError.value = ''
  const isEmailValid = validateEmail(login.email, loginErr)
const isPasswordValid = validatePassword(login.password, loginErr)

if (!isEmailValid || !isPasswordValid) return

  loginLoading.value = true
  try {
    await loginUser({
      email: login.email,
      password: login.password
    })
    close()
  } catch (error) {
    loginApiError.value = error.message || 'Login failed.'
  } finally {
    loginLoading.value = false
  }
}

async function handleSignup() {
  signupApiError.value = ''
  signupApiSuccess.value = ''

  const isEmailValid = validateEmail(signup.email, signupErr)
const isPasswordValid = validatePassword(signup.password, signupErr)

  signupErr.firstName = ''
  if (!signup.firstName.trim()) {
    signupErr.firstName = 'First name is required.'
  }

  signupErr.lastName = ''
  if (!signup.lastName.trim()) {
    signupErr.lastName = 'Last name is required.'
  }

  signupErr.role = ''
  if (!signup.role) {
    signupErr.role = 'Role is required.'
  }

  signupErr.confirmPassword = ''
  if (!signup.confirmPassword) {
    signupErr.confirmPassword = 'Please confirm your password.'
  } else if (signup.password !== signup.confirmPassword) {
    signupErr.confirmPassword = 'Passwords do not match.'
  }

  if (
  !isEmailValid ||
  !isPasswordValid ||
  signupErr.firstName ||
  signupErr.lastName ||
  signupErr.role ||
  signupErr.confirmPassword
) return

  signupLoading.value = true
  try {
    const response = await signupUser({
  firstName: signup.firstName,
  lastName: signup.lastName,
  email: signup.email,
  password: signup.password,
  role: signup.role
})

    signupApiSuccess.value = response?.message || 'Account created successfully.'

    Object.assign(signup, {
      firstName: '',
      lastName: '',
      email: '',
      role: '',
      password: '',
      confirmPassword: ''
    })
    showConfirm.value = false
  } catch (error) {
    signupApiError.value = error.message || 'Signup failed.'
  } finally {
    signupLoading.value = false
  }
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
  background: rgba(18, 14, 35, 0.95);
  border: 1px solid rgba(255, 255, 255, 0.09);
  border-radius: 20px;
  padding: 1.75rem 2rem 2rem;
  box-shadow:
    0 0 0 1px rgba(168, 85, 247, 0.08),
    0 24px 64px rgba(0, 0, 0, 0.6),
    0 0 80px rgba(124, 58, 237, 0.12);
  overflow: hidden;
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
  z-index: 1;
}
.close-btn:hover { background: rgba(255,255,255,0.09); color: rgba(226,232,240,0.85); }

.track-wrap {
  overflow: hidden;
}

.track {
  display: flex;
  width: 200%;
  transition: transform 0.32s cubic-bezier(0.4, 0, 0.2, 1);
}

.track.show-signup {
  transform: translateX(-50%);
}

.panel {
  width: 50%;
  flex-shrink: 0;
}

.panel-head { margin-bottom: 1.5rem; }

.dialog-title {
  font-size: 1.45rem;
  font-weight: 800;
  color: #fff;
  letter-spacing: -0.03em;
  margin-bottom: 0.25rem;
}

.dialog-sub {
  color: rgba(226,232,240,0.4);
  font-size: 0.83rem;
}

.form { display: flex; flex-direction: column; gap: 1rem; }

.name-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.75rem;
}

.field { display: flex; flex-direction: column; gap: 0.35rem; }

.label {
  font-size: 0.77rem;
  font-weight: 600;
  color: rgba(226,232,240,0.6);
  letter-spacing: 0.02em;
}

.input-wrap { position: relative; display: flex; align-items: center; }

:deep(.input-icon) {
  position: absolute;
  left: 0.8rem;
  color: rgba(226,232,240,0.25);
  pointer-events: none;
  z-index: 1;
}

.input,
.select-input {
  width: 100%;
  background: rgba(255,255,255,0.04);
  border: 1px solid rgba(255,255,255,0.08);
  border-radius: 10px;
  color: #e2e8f0;
  font-size: 0.875rem;
  font-family: inherit;
  padding: 0.6rem 0.85rem 0.6rem 2.4rem;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.select-input {
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
}

.input::placeholder { color: rgba(226,232,240,0.18); }
.input:focus,
.select-input:focus { border-color: rgba(168,85,247,0.5); box-shadow: 0 0 0 3px rgba(168,85,247,0.1); }
.input.error,
.select-input.error { border-color: rgba(239,68,68,0.5); }
.input.error:focus,
.select-input.error:focus { box-shadow: 0 0 0 3px rgba(239,68,68,0.1); }
.input-with-toggle { padding-right: 2.6rem; }

.field-spacer { visibility: hidden; pointer-events: none; }

.eye-btn {
  position: absolute;
  right: 0.7rem;
  background: none;
  border: none;
  color: rgba(226,232,240,0.3);
  cursor: pointer;
  padding: 0.2rem;
  display: flex;
  align-items: center;
  border-radius: 5px;
  transition: color 0.2s;
}
.eye-btn:hover { color: rgba(226,232,240,0.75); }

.field-error { font-size: 0.72rem; color: #f87171; }

.api-message {
  font-size: 0.78rem;
  border-radius: 10px;
  padding: 0.7rem 0.85rem;
  margin-bottom: 0.15rem;
}

.api-error {
  color: #fecaca;
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.25);
}

.api-success {
  color: #bbf7d0;
  background: rgba(34, 197, 94, 0.1);
  border: 1px solid rgba(34, 197, 94, 0.25);
}

.submit-btn {
  margin-top: 0.25rem;
  width: 100%;
  background: linear-gradient(135deg, #7c3aed, #a855f7);
  border: none;
  color: #fff;
  font-size: 0.9rem;
  font-weight: 600;
  font-family: inherit;
  padding: 0.68rem;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0 0 22px rgba(168,85,247,0.35);
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 44px;
}
.submit-btn:hover:not(:disabled) { transform: translateY(-1px); box-shadow: 0 0 34px rgba(168,85,247,0.55); }
.submit-btn:disabled { opacity: 0.7; cursor: not-allowed; }

.spinner {
  width: 17px; height: 17px;
  border: 2px solid rgba(255,255,255,0.3);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }

.switch-hint {
  text-align: center;
  margin-top: 1.2rem;
  font-size: 0.8rem;
  color: rgba(226,232,240,0.38);
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

.fade-enter-active, .fade-leave-active { transition: opacity 0.22s ease; }
.fade-enter-active .dialog, .fade-leave-active .dialog { transition: opacity 0.22s ease, transform 0.22s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
.fade-enter-from .dialog, .fade-leave-to .dialog { opacity: 0; transform: scale(0.96) translateY(8px); }

@media (max-width: 520px) {
  .name-row {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
}
</style>