<template>
  <div class="test-connect">
    <h2>バックエンド接続テスト</h2>

    <div class="input-box">
      <input v-model="name" placeholder="お名前を入力してください" />
      <button @click="sendRequest">送信</button>
    </div>

    <div class="message-box">
      <p>バックエンドからのメッセージ: {{ message }}</p>
    </div>
  </div>
</template>

<script>
import { ref } from "vue";

export default {
  name: "TestConnect",
  setup() {
    const name = ref("");
    const message = ref("未接続");

    const sendRequest = async () => {
      if (!name.value.trim()) {
        message.value = "名前を入力してください";
        return;
      }

      try {
        const response = await fetch("http://localhost:8080/api/greet", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            // name.valueの値（"John(これは例)"）を"name"というキーを持つJavaScriptオブジェクトに設定
            // その後、そのオブジェクトをJSON文字列に変換
            //要はオブジェクトや変数(文字列)をJSON文字列({name:"John"})に変換している
            name: name.value,
          }),
        });
        // レスポンスのボディをJSONとして解析
        const data = await response.json();
        // data = { message: "Hello, John!" }
        message.value = data.message;
      } catch (error) {
        console.error("接続エラー:", error);
        message.value = "接続エラーが発生しました";
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

<!-- 写経する際にcssは必要ない -->
<style scoped>
.test-connect {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  text-align: center;
}

.input-box {
  margin: 20px 0;
  display: flex;
  gap: 10px;
  justify-content: center;
}

input {
  padding: 8px;
  width: 200px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.message-box {
  margin: 20px 0;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  padding: 8px 16px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #45a049;
}
</style>
