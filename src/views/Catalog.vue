<!-- <template>
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
      </v-col>
    </v-row>
    <v-btn
      color="primary"
      variant="elevated">

    Add more product
  </v-btn>
    </div>
</template>
<script>
  import { defineComponent } from "vue";
  import ProductItem from "@/components/ProductItem.vue";
  export default defineComponent({
    name: 'Catalog',
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
    router.push({ name: 'ProductDetailUser', params: { id } })
  }
  const goToAdminPage = () => {
    
    router.push({ name: 'EditView' })
  }

 const removeFromProduct = (id) => {
    store.removeFromProduct(id)
  }

  // const addProduct = () => {
  //   store.addProduct()
  //   goToAdminPage(store.size)
  // }

  onMounted(async () => {
    await store.fetchProductsFromDB()
  })
</script>

<style scoped>

</style> -->

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
      </v-col>
    </v-row>           
    </div>
      <v-btn
      @click="goToCatelogView"
      color="primary"
      variant="elevated">
    Go to CatalogAdminView
  </v-btn>
</template>

<script>
  import { defineComponent } from "vue";
  import ProductItem from "@/components/ProductItem.vue";
  export default defineComponent({
    name: 'Catalog',
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
    router.push({ name: 'ProductViewUser', params: { id } })
  }

  const goToCatelogView = () => {
    
    router.push({ name: 'CatalogAdmin' })
  }


 const removeFromProduct = (id) => {
    store.removeFromProduct(id)
  }

 const addProduct = (id) => {
    store.removeFromProduct(id)
  }

  onMounted(async () => {
    await store.fetchProductsFromDB()
  })
</script>

<style scoped>

</style>