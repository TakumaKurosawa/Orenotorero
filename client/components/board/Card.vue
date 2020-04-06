<template>
  <div>
    <v-card raised class="ma-3">
      <v-card-text
        class="text-truncate font-weight-bold subtitle-1"
        @click.stop="dialog = true"
      >
        {{ card.title }}
      </v-card-text>
    </v-card>
    <v-dialog v-model="dialog" max-width="500">
      <v-card>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-icon @click="dialog = false">mdi-close</v-icon>
        </v-card-actions>
        <v-card-title class="headline">
          <span v-if="isEdit"
            ><v-text-field v-model="cardTitle" type="text"
          /></span>
          <span v-else>{{ card.title }}</span>
          <a v-if="isEdit" @click="updateCardTitle"
            ><v-icon>mdi-send</v-icon></a
          >
          <a v-if="isEdit" @click="toggleIsEdit"
            ><v-icon>mdi-pencil-remove</v-icon></a
          >
          <a v-else @click="toggleIsEdit"
            ><v-icon>mdi-square-edit-outline</v-icon></a
          >
        </v-card-title>
        <v-card-text>
          期限:
          <v-menu v-model="deadlinePicker">
            <template v-slot:activator="{ on }">
              <v-btn outlined v-on="on">
                <template v-if="deadline">{{
                  new Date(deadline).toLocaleDateString()
                }}</template>
                <template v-else-if="!card.dead_line">なし</template>
                <template v-else>
                  {{ new Date(card.dead_line).toLocaleDateString() }}
                </template>
              </v-btn>
            </template>
            <v-date-picker
              v-model="deadline"
              @click:date="updateDeadline()"
            ></v-date-picker>
          </v-menu>
        </v-card-text>
        <v-card-title>
          説明:<v-btn
            v-if="!isDescribeEdit"
            @click="isDescribeEdit = true"
            color="blue-grey"
            outlined
            >編集</v-btn
          >
        </v-card-title>
        <v-card-text v-if="isDescribeEdit">
          <v-textarea v-model="describe" outlined auto-grow></v-textarea>
          <v-btn color="success" @click="updateDescribe">保存</v-btn>
          <v-btn icon @click="isDescribeEdit = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-card-text>
        <v-card-text v-else>
          <div class="text--primary">
            {{ describe }}
          </div>
        </v-card-text>
        <v-card-actions>
          <v-btn color="error" @click="deleteCard">削除</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import { Prop } from '~/node_modules/nuxt-property-decorator'

@Component
export default class Card extends Vue {
  dialog = false
  deadline = ''
  deadlinePicker = false
  describe = ''
  isDescribeEdit = false
  isEdit = false
  cardTitle = ''

  created() {
    this.cardTitle = this.card.title
    this.describe = this.card.describe
  }

  @Prop({ type: Object, required: true })
  card!: object

  @Prop({ type: Number, required: true })
  cardIndex!: number

  async updateDeadline() {
    const payload = {
      deadline: this.deadline + ' 00:00:00',
      id: this.card.id
    }
    console.log(payload)
    await this.$axios
      .put('/card/deadline', payload, {
        headers: {
          Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
        }
      })
      .then((res: any) => {
        console.log(res.data)
      })
      .catch((err: any) => {
        console.log(err)
      })
  }

  async updateDescribe() {
    const payload = {
      describe: this.describe,
      id: this.card.id
    }
    console.log(payload)
    await this.$axios
      .put('/card/describe', payload, {
        headers: {
          Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
        }
      })
      .then((res: any) => {
        console.log(res.data)
        this.$store.dispatch('board/getBoardData', {
          boardId: this.$route.params.id,
          token: this.$store.getters['auth/getAuthToken']
        })
        this.isDescribeEdit = false
      })
      .catch((err: any) => {
        console.log(err)
      })
  }

  async deleteCard() {
    const options = {
      headers: {
        Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
      },
      data: {
        id: this.card.id
      }
    }
    console.log(options)
    await this.$axios
      .delete('/card', options)
      .then(() => {
        this.dialog = false
        this.$store.dispatch('board/getBoardData', {
          boardId: this.$route.params.id,
          token: this.$store.getters['auth/getAuthToken']
        })
      })
      .catch((err: any) => {
        console.log(err)
      })
  }

  toggleIsEdit(): void {
    this.isEdit = !this.isEdit
  }

  updateCardTitle(): void {
    const payload = {
      id: this.card.id,
      title: this.cardTitle,
      token: this.$store.getters['auth/getAuthToken']
    }

    console.log(payload)
    this.$store.dispatch('board/updateCardTitle', payload)
  }
}
</script>
