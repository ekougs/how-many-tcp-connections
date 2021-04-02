import ws from 'k6/ws';
import {check} from 'k6';

const host = __ENV.HOST || '127.0.0.1';
const port = __ENV.PORT || 8080;
const targetVUs = __ENV.VUS || 10;
const rampUp = __ENV.RAMP_UP || '5m';
const rampDown = __ENV.RAMP_DOWN || rampUp;
const duration = __ENV.DURATION || '10m';

export let options = {
    stages: [
        { duration: rampUp, target: targetVUs }, // simulate ramp-up of traffic from 1 to target nb of users.
        { duration: duration, target: targetVUs }, // stay at target nb of users for the specified duration
        { duration: rampDown, target: 0 }, // ramp-down to 0 users
    ],
    thresholds: {
        http_req_duration: ['p(99)<1500'], // 99% of requests must complete below 1.5s
    },
};

// From MDN - https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/random
function getRandomIntInclusive(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1) + min); //The maximum is inclusive and the minimum is inclusive
}

function generateConversation(socket) {
    const nbMessages = getRandomIntInclusive(4, 10);
    const sendMessage = (intervalMs) => {
        socket.send(`Message generated to be sent after ${intervalMs} ms`);
    };

    let interval = 0;
    // Generate the n - 1 first messages to be sent async
    for(let i = 0 ; i < nbMessages - 1 ; i++) {
        interval += getRandomIntInclusive(200, 500);
        // If interval is not deeply copied when the function is executed, it will always get its last value
        const currentInterval = JSON.parse(JSON.stringify(interval));
        socket.setTimeout(() => {
            sendMessage(currentInterval);
        }, interval);
    }
    // Generate the last message to be sent async and then close the socket
    interval += getRandomIntInclusive(200, 500);
    socket.setTimeout(() => {
        sendMessage(interval);
        socket.close()
    }, interval);
}

export default () => {
    const url = `ws://${host}:${port}/echo`;
    const params = { tags: { my_tag: 'hello' } };
    const res = ws.connect(url, params, function (socket) {
        socket.on('open', () => {
            console.log('connected');
            generateConversation(socket);
        });
        socket.on('message', (data) => {
            console.log('Message received: ', data);
        });
        socket.on('close', () => console.log('disconnected'));
    });
    check(res, { 'status is 101': (r) => r && r.status === 101 });
}