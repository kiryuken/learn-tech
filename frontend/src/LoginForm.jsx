import { useState } from 'react'
import './LoginForm.css'

/**
 * [DOCUMENTATION: Component LoginForm]
 * Komponen autentikasi (Login & Sign Up) terintegrasi dengan endpoint backend Go:
 * - POST /auth/login   {"email":"...", "password":"..."}
 * - POST /auth/signup  {"email":"...", "password":"..."}
 */
function LoginForm() {
  // [TAG: STATE MANAGEMENT - MODE & INPUTS]
  const [mode, setMode] = useState('login') // 'login' | 'signup'
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  // [TAG: STATE MANAGEMENT - STATUS & RESPON]
  const [statusMessage, setStatusMessage] = useState({ type: '', text: '' })
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [responseData, setResponseData] = useState(null)

  // [TAG: EVENT HANDLER - FORM SUBMIT]
  const handleSubmit = async (e) => {
    e.preventDefault()
    setResponseData(null)

    // [TAG: VALIDATION] - Validasi sederhana input tidak boleh kosong
    if (!email.trim() || !password.trim()) {
      setStatusMessage({
        type: 'error',
        text: 'Email dan password wajib diisi!'
      })
      return
    }

    setIsSubmitting(true)
    setStatusMessage({ type: '', text: '' })

    // [TAG: API ENDPOINT SELECTION] - Menentukan endpoint /auth/login atau /auth/signup
    const endpoint = mode === 'login' ? '/auth/login' : '/auth/signup'

    try {
      // [TAG: API CALL - POST REQUEST]
      const response = await fetch(endpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      const data = await response.json()

      if (!response.ok) {
        // Handling error response dari backend
        throw new Error(data.error || 'Terjadi kesalahan pada server')
      }

      // [TAG: SUCCESS HANDLING]
      setResponseData(data)
      setStatusMessage({
        type: 'success',
        text: mode === 'login'
          ? 'Login berhasil!'
          : `Akun ${data.email} berhasil didaftarkan!`
      })

      if (mode === 'login') {
        setPassword('')
      }
    } catch (err) {
      // [TAG: ERROR HANDLING]
      setStatusMessage({
        type: 'error',
        text: err.message
      })
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    // [TAG: UI CONTAINER] - Container utama form monokrom
    <div className="login-card">
      {/* [TAG: TAB SWITCHER] - Navigasi antar mode Login dan Sign Up */}
      <div className="tab-switcher">
        <button
          type="button"
          className={`tab-btn ${mode === 'login' ? 'active' : ''}`}
          onClick={() => {
            setMode('login')
            setStatusMessage({ type: '', text: '' })
            setResponseData(null)
          }}
        >
          Masuk (Login)
        </button>
        <button
          type="button"
          className={`tab-btn ${mode === 'signup' ? 'active' : ''}`}
          onClick={() => {
            setMode('signup')
            setStatusMessage({ type: '', text: '' })
            setResponseData(null)
          }}
        >
          Daftar (Sign Up)
        </button>
      </div>

      <div className="login-header">
        <h2 className="login-title">
          {mode === 'login' ? 'Masuk ke Akun' : 'Buat Akun Baru'}
        </h2>
        <p className="login-subtitle">
          {mode === 'login'
            ? 'Masukkan email dan password untuk masuk'
            : 'Daftarkan email dan password baru'}
        </p>
      </div>

      {/* [TAG: ALERT MESSAGE] - Notifikasi sukses / error */}
      {statusMessage.text && (
        <div className={`login-alert login-alert-${statusMessage.type}`}>
          {statusMessage.text}
        </div>
      )}

      {/* [TAG: FORM ELEMENT] */}
      <form onSubmit={handleSubmit} className="login-form">
        {/* [TAG: INPUT FIELD - EMAIL] */}
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input
            type="email"
            id="email"
            name="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="dev@example.com"
            autoComplete="email"
            required
          />
        </div>

        {/* [TAG: INPUT FIELD - PASSWORD] */}
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="password123"
            autoComplete={mode === 'login' ? 'current-password' : 'new-password'}
            required
          />
        </div>

        {/* [TAG: SUBMIT BUTTON] */}
        <button
          type="submit"
          className="login-button"
          disabled={isSubmitting}
        >
          {isSubmitting
            ? 'Memproses...'
            : mode === 'login'
              ? 'Masuk'
              : 'Daftar'}
        </button>
      </form>

      {/* [TAG: JSON RESPONSE BOX] - Menampilkan balasan JSON dari server */}
      {responseData && (
        <div className="response-box">
          <div className="response-title">Respon Server (JSON):</div>
          <pre className="response-json">
            {JSON.stringify(responseData, null, 2)}
          </pre>
        </div>
      )}
    </div>
  )
}

export default LoginForm
