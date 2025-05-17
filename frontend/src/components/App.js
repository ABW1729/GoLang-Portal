import React, { useState } from 'react';
import ReceptionistPortal from './ReceptionistPortal';
import DoctorPortal from './DoctorPortal';

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [token, setToken] = useState('');
  const [role, setRole] = useState('');
  const [error, setError] = useState('');

  const login = async () => {
    const res = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    });

    const data = await res.json();
    if (data.token) {
      setToken(data.token);
      setRole(data.role);
    } else {
      setError(data.error || 'Login failed');
    }
  };

  const logout = () => {
    setToken('');
    setRole('');
  };

  if (token && role === 'receptionist') {
    return <ReceptionistPortal token={token} logout={logout} />;
  }

  if (token && role === 'doctor') {
    return <DoctorPortal token={token} logout={logout} />;
  }

  return (
    <div className="container">
      <h1>Login</h1>
      <input placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button className="primary" onClick={login}>Login</button>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
}

export default App;
