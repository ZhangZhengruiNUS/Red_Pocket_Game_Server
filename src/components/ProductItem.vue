<template>
  <v-sheet class="ma-2 pa-2">
    <v-card
        class="product"
    >
      <v-img
          :src="productData.thumbnail"
          height="200px"
          cover
      />

      <v-card-title>
        {{ productData.brand }}
      </v-card-title>

      <v-card-subtitle>
        $ {{ productData.price }}
      </v-card-subtitle>

      <v-card-text>
        {{ productData.description }}
      </v-card-text>

      <v-card-actions>
        <v-btn @click="goToProductPage(productData.id)">
          GO TO VIEW
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-sheet>
</template>

<script>
import { ref, defineProps } from "vue";

export default {
  name: 'ProductItem',
  props: {
    productData: {
      type: Object,
      required: true,
    },
  },
  setup(props, { emit }) {
    const goToProductPage = (productId) => {
      emit('item-clicked', productId);
    };

    return {
      goToProductPage,
    };
  },
};
</script>