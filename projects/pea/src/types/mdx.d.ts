declare module '*.mdx' {
  import type { ComponentType } from 'react'

  interface MDXProps {
    components?: Record<string, React.ComponentType<any>>
  }

  const MDXComponent: ComponentType<MDXProps>

  export default MDXComponent
}
