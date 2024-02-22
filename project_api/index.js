const express = require('express');
const { Pool } = require('pg');

const pool = new Pool({
    user: 'gusatvowiller',
    host: '',
    database: '',
    password: '1234',
    port: 5432,
});

const app = express();
const port = 3000;

app.get('/tb01', async (req, res) => {
    try {
        const result = await pool.query('SELECT * FROM TB01 ORDER BY col_dt DESC LIMIT 10');
        res.json(result.rows);
    } catch (err) {
        console.error('Erro ao executar a consulta', err);
        res.status(500).json({ error: 'Erro ao executar a consulta' });
    }
});

app.listen(port, () => {
    console.log(`Servidor rodando na porta ${port}`);
});
