<template>
  <v-sheet class="ma-2 pa-2">
    <v-card
        class="productUser"
    >
      <v-img
          :src="props.productData.thumbnail"
          height="200px"
          cover
      />

      <v-card-title>
        {{ props.productData.brand }}
      </v-card-title>

      <v-card-subtitle>
        $ {{ props.productData.price }}
      </v-card-subtitle>

      <v-card-text>
        {{ props.productData.description }}
      </v-card-text>

      <v-card-actions>
        <v-btn @click="goToProductPage(props.productData.id)">
          Redeem
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-sheet>
</template>

<script>
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: 'ProductItemUser',
  props: {
    productData: {
      type: Object,
      required: true,
    },
  },
  setup(props, context) {
    const goToProductPage = (productId) => {
      context.emit('item-clicked', productId);
    };

    // Return variables and functions you want to expose in the template
    return {
      goToProductPage,
    };
  },
});
</script>
