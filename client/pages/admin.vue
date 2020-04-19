<template>
  <v-layout
    column
    justify-center
    align-center
  >
    <v-flex
      xs12
      sm8
      md6
    >
      <v-card>
        <v-card-title>アカウント管理画面</v-card-title>
        <v-card-text>
          <button @click="logout">Logout</button>
        </v-card-text>
      </v-card>
      <br>
      <br>
      <v-card>
        <v-card-title>登録フォーム管理</v-card-title>
        <v-card-text>
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
        </v-card-text>
      </v-card>
      <br>
      <br>
      <v-card>
        <v-card-title>アカウント管理</v-card-title>
        <v-card-text>
          <v-alert type="error" v-if="account_error">
            {{ account_error }}
          </v-alert>
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
        </v-card-text>
      </v-card>
      <br>
      <br>
      <v-card>
        <v-card-title>アカウントアクティベイト</v-card-title>
        <v-card-text>
          <v-checkbox v-model="autoActivate" label="自動アクティベイト"/>
          <v-alert type="error" v-if="activate_error">
            {{ activate_error }}
          </v-alert>
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
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  middleware: 'auth',
  data() {
    return {
      new_form_time: 600,
      form_url: "",
      form_time: 0,
      form_error: "",
      account_error: "",
      activate_error: "",
      required: value => !!value || "必須項目です",
      number_check: value => !isNaN(value) || "半角英数字のみで入力してください",
      autoActivate: false,

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
  head() {
    return {
      title: 'アカウント管理画面',
    }
  },
  mounted() {
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
        this.accounts_desserts = result
      })
      .catch((error) => {
        this.account_error =  error.response.data.Msg
      })
    setInterval(() => {
      this.$axios.$post('/api/admin/get_all_accounts')
        .then((result) => {
          this.accounts_desserts = result
        })
    }, 5000)

    this.$axios.$post('/api/admin/get_all_non_active_account_id')
      .then((result) => {
        this.activate_desserts = result
      })
      .catch((error) => {
        this.activate_error = error.response.data.Msg
      })
    setInterval(() => {
      this.$axios.$post('/api/admin/get_all_non_active_account_id')
        .then((result) => {
          this.activate_desserts = result
          if (!this.autoActivate) {
            return
          }
          this.activate_desserts.forEach(account => {
            const params = new URLSearchParams()
            params.append('ID', account.ID)
            this.$axios.$post('/api/admin/activation', params)
              .then((result) => {
              })
          })
        })
    }, 5000)

  },
  methods: {
    logout() {
      this.$auth.logout()
    },
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
          this.activate_error = ''
          const index = this.activate_desserts.indexOf(item)
          this.activate_desserts.splice(index, 1)
        })
        .catch((error) => {
          this.activate_error =  error.response.data.Msg
        })
    },
    delete_account(account) {
      const params = new URLSearchParams()
      params.append('ID', account.ID)
      this.$axios.$post('/api/admin/delete_account', params)
        .then(() => {
          this.account_error = ''
          const index = this.accounts_desserts.indexOf(account)
          this.accounts_desserts.splice(index, 1)
        })
        .catch((error) => {
          this.account_error =  error.response.data.Msg
        })

    }
  }
}
</script>