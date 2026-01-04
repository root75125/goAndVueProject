<template>
  <div>
    <input type="text" v-model="nameV" placeholder="pls w name" />
    <button @click="sendRequest">send</button>
  </div>
  <div>
    <label>{{ message }}</label>
  </div>
</template>

<script>
import { ref } from "vue";
export default {
  setup() {
    const nameV = ref("");
    const message = ref("unconnect");
    const sendRequest = async () => {
      if (!nameV.value.trim()) {
        message.value = "pls w name";
      }
      try {
        const response = await fetch("http://localhost:8080/api/greet", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            name: nameV.value,
          }),
        });
        const data = await response.json();
        message.value = data.message;
      } catch (erorr) {
        console.error("error:", erorr);
        message.value = "error happen";
      }
    };
    return {
      nameV,
      message,
      sendRequest,
    };
  },
};
</script>

<style lang="scss" scoped></style>
