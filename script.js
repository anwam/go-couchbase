import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
    const age = Math.random() * 100
    const payload = JSON.stringify({
        name: "billy",
        age: age,
        city: "korat"
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post('http://localhost:3000/user', payload, params);
}