<template>
  <div>
    <LabelIcon
      :text-label="'パーソナルボード'"
      :icon-name="'mdi-account-outline'"
    ></LabelIcon>
    <v-row>
      <Card
        v-for="item in boardData"
        :key="item.id"
        :card-title="item.title"
        :card-height="100"
        :card-width="180"
        :card-img="item.bg_img"
        @action="transitionBoard(item.id)"
      ></Card>
      <Card
        :card-title="'ボードを追加'"
        :card-height="100"
        :card-width="180"
        :card-color="'Gray'"
        @action="validDialog()"
      ></Card>
    </v-row>
    <v-row justify="center">
      <v-dialog v-model="dialog">
        <v-card>
          <v-card-title>
            <span class="headline">新規ボード作成</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                    v-model="boardTitle"
                    label="ボードタイトル"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12">
                  <v-text-field
                    v-model="imageURL"
                    label="画像URL"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="createBoard(boardTitle, imageURL)"
              >作成</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </div>
</template>
<script lang="ts">
import { Vue, Prop, Component } from 'nuxt-property-decorator'
import LabelIcon from '@/components/atom/LabelIcon.vue'
import Card from '@/components/atom/Card.vue'

@Component({
  components: {
    LabelIcon,
    Card
  }
})
export default class PersonalBoardList extends Vue {
  dialog = false
  boardTitle = ''
  imageURL = ''
  async createBoard(title: string, img: string) {
    this.dialog = false
    const payload = {
      title,
      img
    }
    await this.$axios
      .post('/board', payload, {
        headers: {
          Authorization: 'Bearer ' + this.$store.getters['auth/getAuthToken']
        }
      })
      .then((res: any) => {
        console.log(res.data)
        this.$emit('action')
      })
      .catch((err: any) => {
        console.log(err)
        this.boardTitle = ''
        this.imageURL = ''
      })
  }

  transitionBoard(id: Number) {
    this.$router.push('/board/' + id)
  }

  validDialog() {
    this.dialog = true
  }

  @Prop({ type: Array, required: true })
  boardData!: Array<string>
}
</script>
