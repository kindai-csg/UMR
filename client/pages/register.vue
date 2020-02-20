<template>
  <v-card>
    <v-card-title>近畿大学電子計算機研究会 入部フォーム</v-card-title>
    <v-card-text>
      <v-form ref="register_form">
        <v-text-field
          v-model="id"
          label="ユーザーID (4文字~15文字, 英数字)"
          :rules="[required, id_length, alpha_num_check]"
        />
        <v-text-field
          v-model="password"
          label="パスワード (8文字以上)"
          type="password"
          :rules="[required, password_length]"
        />
        <v-text-field
          v-model="passwordconfirm"
          label="パスワード (確認)"
          type="password"
          :rules="[required, password_confirm]"
        />
        <v-text-field
          v-model="studentnumber"
          label="学籍番号 (例 2010370300)"
          :rules="[required, number_check]"
        />
        <v-text-field
          v-model="name"
          label="氏名"
          :rules="[required]"
        />
        <v-text-field
          v-model="mailaddress"
          label="メールアドレス"
          :rules="[required, mail_check]"
        />
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-btn text v-on:click="submit">送信する</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data() {
      return {
        id: "",
        password: "",
        passwordconfirm: "",
        studentnumber: "",
        name: "",
        mailaddress: "",
        required: value => !!value || "必須項目です",
        id_length: value => (value.length >= 4 && value.length <= 15) || "4文字以上15文字以内で入力してください",
        password_length: value => value.length >= 8 || "8文字以上で入力してください",
        password_confirm: value => {
          if (this.password !== value) {
            return "パスワードが一致しません"
          }
        },
        number_check: value => !isNaN(value) || "半角数字のみで入力してください",
        alpha_num_check: value => {
          if (!value.match(/^[A-Za-z0-9]*$/))
            return "半角英数字のみで入力してください"
        },
        mail_check: value => {
          if (!value.match(/.+@.+\..+/) || !value.match(/^[A-Za-z0-9@.+]*$/))
            return "メールアドレスを入力してください"
        }
      }
  },
  methods: {
    submit() {
      if (this.$refs.register_form.validate()) {
        console.log("submit")
      } else {
        console.log("out")
      }
    }
  }
}
</script>