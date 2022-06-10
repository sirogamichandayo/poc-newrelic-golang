import { check, fail } from 'k6'
import http from 'k6/http'

const HOST = 'http://localhost:8080'

export default function() {
    inserUsers();
    getUsers();
}

function getUsers() {
    let response = http.get(`${HOST}/users`)

    let responseJson
    try {
        responseJson = response.json()
    } catch (err) {
        console.log({ r: response.body, status: response.status })
    }

    
    check(response, {
        'status is 200': (r) => r.status === 200,
    })
}

function inserUsers() {
    
    const data = {
        name: "test",
        user_name: `${String(new Date().getTime())}-${(Math.random() + 1).toString(36).substring(7)}`,
        age: 21,
        email: "test@teste.com"
    }
    let response = http.post(`${HOST}/users/new`, JSON.stringify(data), {
        headers: { 'Content-Type': 'application/json' }
    })
    
    check(response, {
        'status is 201': (r) => r.status === 201,
    })
}