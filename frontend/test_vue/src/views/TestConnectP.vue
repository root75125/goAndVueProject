<template>
    <div>
        <input type="text" v-model="namev" placeholder="pls w n" />
        <button @click="sendRequest">send</button>
    </div>
    <div>
        <label>{{ message }}</label>
    </div>
</template>

<script>
import { ref } from 'vue';

export default {
    setup() {
        const namev = ref("");
        const message = ref("unconnect");
        const sendRequest = async () => {
            if (!namev.value.trim()) {
                message.value = "pls w n"
                return
            }
            try {
                const response = await fetch("http://localhost:8080/api/greet", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        name: namev.value
                    })
                })
                const data = await response.json()
                message.value = data.message
            } catch (error) {
                console.error("error:", error)
                message.value = "error happen"
                return
            }
        }
        return {
            namev,
            message,
            sendRequest,
        };
    },
};
</script>
