const express = require('express');
const { Pool } = require('pg');
const cors = require("cors")

const pool = new Pool({
    user: 'postgres',
    host: 'localhost',
    database: 'postgres',
    password: '1234',
    port: 5432,
});

const app = express();
const port = 3000;

app.use(cors({origin: (origin, cb) => {
    cb(null, origin);
}, }))

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
