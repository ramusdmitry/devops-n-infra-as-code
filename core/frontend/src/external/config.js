const CHAT_PORT = "8040"
const POST_PORT = "8036"
const AUTH_PORT = "8000"
const GROUP_PORT = "8050"
const COMM_PORT = "8032"
const HOST = "localhost"

export const CHAT_URL = `ws://${HOST}:${CHAT_PORT}/room`
export const POST_URL = `http://${HOST}:${POST_PORT}/api/posts`
export const AUTH_URL = `http://${HOST}:${AUTH_PORT}/api/auth`
export const COMMS_URL = `http://${HOST}:${COMM_PORT}/api/comms`

export const GROUP_URL = `http://${HOST}:${GROUP_PORT}/api/group`
