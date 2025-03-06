<script setup>
import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
import ToolingIcon from './icons/IconTooling.vue'
import EcosystemIcon from './icons/IconEcosystem.vue'
import CommunityIcon from './icons/IconCommunity.vue'
import SupportIcon from './icons/IconSupport.vue'
</script>

<template>
  <WelcomeItem>

    <template #icon>
      <CommunityIcon />
    </template>

    <template #heading>Community</template>

    Got stuck? Ask your question on StackOverflow.

  </WelcomeItem>

  <WelcomeItem>

    <template #icon>
      <SupportIcon />
    </template>

    <template #heading>Support Vue</template>

    <template #subheading>Custom Subheading</template>

    You can help us by becoming a sponsor.

  </WelcomeItem>
</template>
