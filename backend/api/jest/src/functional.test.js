const axios = require('axios');

const host = process.env['BACKEND_HOST'];
console.log(`configured host: ${host}`)

const client = axios.create({
    baseURL: 'https://some-domain.com/api/',
    timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});

// client.get('/')
//     .then(response => {
//         console.log('success response: ');
//         console.log(response);
//     })
//     .catch(err => {
//         console.log('error response: ');
//         console.log(err);
//     });

describe('empty database', () => {
    client.get('/')
        .then(response => {
            console.log('success response: ');
            console.log(response);
        })
        .catch(err => {
            console.log('error response: ');
            console.log(err);
        });

    test('returns no records', () => {
        expect(1).toBe(1);
    });
});