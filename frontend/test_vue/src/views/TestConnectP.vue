<template>
  <div>
    <input type="text" v-model="name" placeholder="pls w n" />
    <button @click="sendRequest">send</button>
  </div>
  <label>{{ message }}</label>
</template>

<script>
import { ref } from "vue";
export default {
  setup() {
    const name = ref("");
    const message = ref("unconnect");
    const sendRequest = async () => {
      if (name.value == "") {
        message.value = "pls w name";
        return;
      }
      try {
        const response = await fetch("http://localhost:8080/api/greet", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            name: name.value,
          }),
        });
        const data = await response.json();
        message.value = data.message;
      } catch (error) {
        console.error("error:", error);
        message.value = "happen error";
      }
    };
    return {
      name,
      message,
      sendRequest,
    };
  },
};
</script>

<style lang="scss" scoped></style>
