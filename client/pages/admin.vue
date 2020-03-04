<template>
  <v-expansion-panels>
    <v-expansion-panel>
      <v-expansion-panel-header>登録フォーム管理</v-expansion-panel-header>
      <v-expansion-panel-content>
        <v-alert type="error" v-if="form_error">
          {{ form_error }}
        </v-alert>
        <div v-if="form_url&&form_time">
          フォームURL: {{ form_url }}
          <br>
          残り時間: {{ form_time }} 秒 
        </div>
        <br>
        <v-form ref="create_form">
          <v-text-field
          v-model="new_form_time"
          label="有効期限(秒)"
          :rules="[required, number_check]"
          />
          <v-btn text v-on:click="create_form">作成</v-btn>
        </v-form>
      </v-expansion-panel-content>
    </v-expansion-panel>
    <v-expansion-panel>
      <v-expansion-panel-header>アカウント管理</v-expansion-panel-header>
      <v-expansion-panel-content>
        アカウント一覧です.
      </v-expansion-panel-content>
    </v-expansion-panel>
    <v-expansion-panel>
      <v-expansion-panel-header>アカウントアクティベイト</v-expansion-panel-header>
      <v-expansion-panel-content>
        アカウントを有効にします.
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script>
export default {
  data() {
    return {
      new_form_time: 600,
      form_url: "",
      form_time: 0,
      form_error: "",
      required: value => !!value || "必須科目です",
      number_check: value => !isNaN(value) || "半角英数字のみで入力してください"
    }
  },
  created() {
    this.$axios.$post('/api/admin/get_register_form')
      .then((result) => {
        this.form_url = location.origin + "/register?token=" + result.Token
        this.form_time = result.Time
      })
      setInterval(() => {
        if (this.form_time > 0) {
          this.form_time -= 1
        }
      }, 1000)
  },
  methods: {
    create_form() {
      if (this.$refs.create_form.validate()) {
        const params = new URLSearchParams()
        params.append('Time', this.new_form_time)
        this.$axios.$post('/api/admin/create_register_form', params)
          .then((result) => {
            this.form_url = location.origin + "/register?token=" + result.Token
            this.form_time = result.Time
          })
          .catch((e) => {
            if (e.response) {
              this.form_error = e.response.data.Msg
            } else {
              this.form_error = "予期せぬなエラーが発生しました." 
            }
          })
      }
    }
  }
}
</script>