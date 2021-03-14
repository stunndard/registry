<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="stunndard.registry.registry" type="name" :fields="[ ['creator', 1, 'string'] , ['owner', 2, 'string'] , ['price', 3, 'string'] , ['onsale', 4, 'bool'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
