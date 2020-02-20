<template>
  <v-card>
    <v-card-title>確認コード入力</v-card-title>
    <v-card-text>
      <v-form ref="code_form">
        <v-text-field
          v-model="code"
          label="確認コードを入力"
          :rules="[required, id_length, alpha_num_check]"
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
      code: "",
      required: value => !!value || "必須項目です"
    }
  },
  methods: {
    submit() {
      if (this.$refs.code_form.validate()) {
        console.log("submit")
      } else {
        console.log("out")
      }
    }
  }
}
</script>