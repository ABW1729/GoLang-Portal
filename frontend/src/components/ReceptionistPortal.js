import React, { useState, useEffect } from 'react';

function ReceptionistPortal({ token, logout }) {
  const [patients, setPatients] = useState([]);
  const [form, setForm] = useState({ name: '', age: '', gender: '' });
  const [isEditing, setIsEditing] = useState(false);
  const [editId, setEditId] = useState(null);

  const fetchPatients = async () => {
    try {
      const res = await fetch('http://localhost:8080/api/receptionist/patients', {
        headers: { Authorization: `Bearer ${token}` },
      });
      const data = await res.json();
      setPatients(data);
    } catch (err) {
      console.error("Error fetching patients", err);
    }
  };

  const handleInput = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async () => {
    const url = isEditing
      ? `http://localhost:8080/api/receptionist/patients/${editId}`
      : 'http://localhost:8080/api/receptionist/patients';

    const method = isEditing ? 'PUT' : 'POST';

    await fetch(url, {
      method,
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        ...form,
        age: parseInt(form.age, 10)
      }),
    });

    setForm({ name: '', age: '', gender: '' });
    setIsEditing(false);
    setEditId(null);
    fetchPatients();
  };

  const handleEdit = (patient) => {
    setForm({ name: patient.name, age: patient.age, gender: patient.gender });
    setIsEditing(true);
    setEditId(patient.ID); // using 'id' as per the mapped value
  };

  const handleDelete = async (id) => {
    await fetch(`http://localhost:8080/api/receptionist/patients/${id}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${token}` },
    });
    fetchPatients();
  };

  useEffect(() => {
    fetchPatients();
  }, []);

  return (
    <div className="container">

  <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '20px' }}>
    <h2>Receptionist Portal</h2>
    <div style={{ display: 'flex'}}>
      <button className="primary" onClick={fetchPatients} style={{ marginRight: '10px' }}>Refresh</button>
      <button className="primary" onClick={logout}>Logout</button>
    </div>
  </div>

      <input
        name="name"
        placeholder="Name"
        value={form.name}
        onChange={handleInput}
      />
      <input
        name="age"
        placeholder="Age"
        value={form.age}
        onChange={handleInput}
      />

 <div style={{ margin: '10px 0', display: 'flex', gap: '20px' }}>
  <label style={{ display: 'flex', alignItems: 'center' }}>
    <input
      type="radio"
      name="gender"
      value="male"
      checked={form.gender === 'male'}
      onChange={handleInput}
    />
    <span style={{ marginLeft: '5px' }}>Male</span>
  </label>
  <label style={{ display: 'flex', alignItems: 'center' }}>
    <input
      type="radio"
      name="gender"
      value="female"
      checked={form.gender === 'female'}
      onChange={handleInput}
    />
    <span style={{ marginLeft: '5px' }}>Female</span>
  </label>
</div>



      <button className="primary" onClick={handleSubmit}>
        {isEditing ? 'Update Patient' : 'Add Patient'}
      </button>

      {patients && patients.length > 0 && (
        <h3 style={{ marginTop: '30px' }}>Patient List</h3>
      )}

 {patients && patients.map((p) => (
  <div key={p.id} className="patient-card" style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
    <div>
      <strong>{p.name}</strong> — {p.age} / {p.gender}
    </div>
    <div style={{ display: 'flex'}}>
      <button className="primary" onClick={() => handleEdit(p)} style={{ marginRight: '8px' }}>Edit</button>
      <button className="primary" onClick={() => handleDelete(p.ID)}>Delete</button>
    </div>
  </div>
))}
    </div>
  );
}

export default ReceptionistPortal;

