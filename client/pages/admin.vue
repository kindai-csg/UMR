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
        <v-data-table
          :headers="accounts_headers"
          :items="accounts_desserts"
          :items-per-page="5"
          class="elevation-1"
        >
          <template v-slot:item.actions="{ item }">
            <v-icon
              small
              class="mr-2"
              @click="delete_account(item)"
            >
              mdi-delete
            </v-icon>
          </template>
        </v-data-table>
      </v-expansion-panel-content>
    </v-expansion-panel>
    <v-expansion-panel>
      <v-expansion-panel-header>アカウントアクティベート</v-expansion-panel-header>
      <v-expansion-panel-content>
        <v-data-table
          :headers="activate_headers"
          :items="activate_desserts"
          :items-per-page="5"
          class="elevation-1"
        >
          <template v-slot:item.actions="{ item }">
            <v-icon
              small
              class="mr-2"
              @click="activate_item(item)"
            >
              mdi-account-check
            </v-icon>
          </template>
        </v-data-table>
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
      number_check: value => !isNaN(value) || "半角英数字のみで入力してください",

      accounts_headers: [
        {
          text: "ユーザーID",
          align: "start",
          value: "ID",
        },
        {
          text: "学籍番号",
          value: "StudentNumber",
        },
        {
          text: "名前",
          value: "Name",
        },
        {
          text: "メールアドレス",
          value: "EmailAddress",
        },
        {
          text: "アクション",
          value: 'actions',
          sortable: false,
        }
      ],
      accounts_desserts: [],
      activate_headers: [
        {
          text: "ユーザーID",
          align: "start",
          value: "ID",
        },
        {
          text: "アクション",
          value: 'actions',
          sortable: false,
        }
      ],
      activate_desserts: [],
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

    this.$axios.$post('/api/admin/get_all_accounts')
      .then((result) => {
        console.log(result)
        this.accounts_desserts = result
      })
      .catch((error) => {
        console.log(error)
      })

    this.$axios.$post('/api/admin/get_all_non_active_account_id')
      .then((result) => {
        console.log(result)
        this.activate_desserts = result
      })
      .catch((error) => {
        console.log(error)
      })
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
    },
    activate_item(item) {
      const params = new URLSearchParams()
      params.append('ID', item.ID)
      this.$axios.$post('/api/admin/activation', params)
        .then(() => {
          console.log("success")
          const index = this.activate_desserts.indexOf(item)
          this.activate_desserts.splice(index, 1)
        })
        .catch((error) => {
          console.log(error)
        })
    },
    delete_account(account) {
      const params = new URLSearchParams()
      console.log(account)
      params.append('ID', account.ID)
      this.$axios.$post('/api/admin/delete_account', params)
        .then(() => {
          const index = this.accounts_desserts.indexOf(account)
          this.accounts_desserts.splice(index, 1)
        })
        .catch((error) => {
          console.log(error)
        })

    }
  }
}
</script>