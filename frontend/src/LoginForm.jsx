import { useState } from 'react'
import './LoginForm.css'

/**
 * [DOCUMENTATION: Component LoginForm]
 */
function LoginForm() {
  // [TAG: STATE MANAGEMENT] - Menyimpan input user dan status form
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [statusMessage, setStatusMessage] = useState({ type: '', text: '' })
  const [isSubmitting, setIsSubmitting] = useState(false)

  // [TAG: EVENT HANDLER] - Menangani proses submit form login
  const handleSubmit = (e) => {
    e.preventDefault()

    // [TAG: VALIDATION] - Validasi sederhana input tidak boleh kosong
    if (!username.trim() || !password.trim()) {
      setStatusMessage({
        type: 'error',
        text: 'Username dan password wajib diisi!'
      })
      return
    }

    // [TAG: SUBMIT SIMULATION] - Simulasi proses pengiriman data login
    setIsSubmitting(true)
    setStatusMessage({ type: '', text: '' })

    setTimeout(() => {
      setIsSubmitting(false)
      setStatusMessage({
        type: 'success',
        text: `Selamat datang, ${username}! Login berhasil.`
      })
      // Reset password setelah login berhasil
      setPassword('')
    }, 1000)
  }

  return (
    // [TAG: UI CONTAINER] - Wrapper utama form login monokrom
    <div className="login-card">
      <div className="login-header">
        <h2 className="login-title">LOGIN</h2>
        <p className="login-subtitle"></p>
      </div>

      {/* [TAG: ALERT MESSAGE] - Display pesan error atau sukses */}
      {statusMessage.text && (
        <div className={`login-alert login-alert-${statusMessage.type}`}>
          {statusMessage.text}
        </div>
      )}

      {/* [TAG: FORM ELEMENT] - Form input login */}
      <form onSubmit={handleSubmit} className="login-form">
        {/* [TAG: INPUT FIELD - USERNAME] */}
        <div className="form-group">
          <label htmlFor="username">Username</label>
          <input
            type="text"
            id="username"
            name="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="Masukkan username"
            autoComplete="username"
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
            placeholder="••••••••"
            autoComplete="current-password"
          />
        </div>

        {/* [TAG: SUBMIT BUTTON] - Tombol aksi login */}
        <button
          type="submit"
          className="login-button"
          disabled={isSubmitting}
        >
          {isSubmitting ? 'Memproses...' : 'Masuk'}
        </button>
      </form>
    </div>
  )
}

export default LoginForm
