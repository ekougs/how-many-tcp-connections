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

export default () => {
    const url = `ws://${host}:${port}/echo`;
    const params = { tags: { my_tag: 'hello' } };
    const res = ws.connect(url, params, function (socket) {
        socket.on('open', () => {
            // TODO send 2-5 messages after 200-500 separated by 500-1000 ms each
            console.log('connected');
            socket.send('Hello');
        });
        socket.on('message', (data) => {
            console.log('Message received: ', data);
            // TODO close after the last message above
            socket.close();
        });
        socket.on('close', () => console.log('disconnected'));
    });
    check(res, { 'status is 101': (r) => r && r.status === 101 });
}