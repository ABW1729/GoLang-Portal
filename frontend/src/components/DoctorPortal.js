import React, { useState, useEffect } from 'react';

function DoctorPortal({ token, logout }) {
  const [patients, setPatients] = useState([]);
  const [editingId, setEditingId] = useState(null);
  const [notes, setNotes] = useState('');

  const fetchPatients = async () => {
    try {
      const res = await fetch('http://localhost:8080/api/doctor/patients', {
        headers: { Authorization: `Bearer ${token}` },
      });
      const data = await res.json();
      setPatients(Array.isArray(data) ? data : []); // safety
    } catch (err) {
      console.error("Error fetching patients", err);
    }
  };

  const startEditing = (patient) => {
    setEditingId(patient.ID || patient.id);
    setNotes(patient.notes || '');
  };

  const updateNotes = async () => {
    await fetch(`http://localhost:8080/api/doctor/patients/${editingId}/notes`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`
      },
      body: JSON.stringify({ notes })
    });
    alert('Notes updated');
    setEditingId(null);
    setNotes('');
    fetchPatients();
  };

  const cancelEditing = () => {
    setEditingId(null);
    setNotes('');
  };

  useEffect(() => {
    fetchPatients();
  }, []);

  return (
    <div className="container">
      <button className="secondary" onClick={logout} style={{ float: 'right', marginBottom: '20px' }}>
        Logout
      </button>
      <h2>Doctor Portal</h2>

      {patients.map((p) => (
        <div key={p.ID || p.id} className="patient-card" style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <div>
            <strong>{p.name}</strong> — {p.age} / {p.gender}
          </div>
          <button
            className="primary"
            onClick={() => startEditing(p)}
            style={{ padding: '6px 12px', fontSize: '16px',width: '100px'  }}
          >
            Edit Notes
          </button>
        </div>
      ))}

      {editingId && (
        <div className="patient-card" style={{ marginTop: '20px' }}>
          <h4>Editing Notes for Patient ID {editingId}</h4>
          <textarea
            style={{ width: '100%', height: 100, marginBottom: '10px' }}
            placeholder="Enter notes..."
            value={notes}
            onChange={(e) => setNotes(e.target.value)}
          />
          <div style={{ display: 'flex', gap: '10px' }}>
            <button className="primary" onClick={updateNotes} style={{ padding: '8px 14px' }}>Save Notes</button>
            <button className="primary" onClick={cancelEditing} style={{ padding: '8px 14px' }}>Cancel</button>
          </div>
        </div>
      )}
    </div>
  );
}

export default DoctorPortal;

