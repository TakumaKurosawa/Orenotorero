<template>
  <div>
    <LabelIcon
      :text-label="'パーソナルボード'"
      :icon-name="'mdi-account-outline'"
    ></LabelIcon>
    <v-row>
      <div v-if="boardData">
        <Card
          v-for="item in boardData"
          :key="item.id"
          :card-title="item.title"
          :card-height="100"
          :card-width="180"
          :card-img="item.bg_img"
        ></Card>
      </div>
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
  @Prop({ type: Array, required: true })
  boardData!: Array<string>
}
</script>
