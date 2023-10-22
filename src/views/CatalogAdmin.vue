<template>
  <div class="products-list">
    <v-text-field clearable label="Label" prepend-icon="$vuetify"></v-text-field>
    <v-row no-gutters>
      <v-col
          v-for="product in store.products"
          :key="product.id"
          cols="12"
          sm="4"
      >
        <product-item
            :product-data="product"
            @item-clicked="goToProductPage"
            @click="goToProductPage(product.id)"
        />
    <button @click="removeFromProduct(product.id)">Remove</button>
      </v-col>
    </v-row>           
    </div>
      <v-btn
      @click="addProduct"
      color="primary"
      variant="elevated">
    Add more
  </v-btn>
  <br><br>
  <v-btn
      @click="goToCatelogView"
      color="primary"
      variant="elevated">
   Go to user page
  </v-btn>
</template>

<script>
  import { defineComponent } from "vue";
  import ProductItem from "@/components/ProductItem.vue";
  export default defineComponent({
    name: 'CatalogAdmin',
    components: {
      ProductItem
    }
  })
</script>

<script setup>
  import { onMounted, ref } from "vue";
  import { productsStore } from "@/stores/products";
  import { useRouter } from "vue-router";

  const store = productsStore()
  const router = useRouter()

  const search = ref('')

  const goToProductPage = (id) => {
    router.push({ name: 'ProductView', params: { id } })
  }

  const goToCatelogView = () => {
    
    router.push({ name: 'Catalog' })
  }


 const removeFromProduct = (id) => {
    store.removeFromProduct(id)
  }

//  const addProduct = (id) => {
//     store.removeFromProduct(id)
//   }

  const addProduct = () => {
    store.addProduct()
    // goToProductPage(store.products.length)
    //goToProductPage(store.size)

  }

  onMounted(async () => {
    await store.fetchProductsFromDB()
  })
</script>

<style scoped>

</style>