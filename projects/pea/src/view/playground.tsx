import MyMdxPage from '&/page/mdx-page.mdx'
import { CustomButton } from '&/view'
import { MDXProvider } from '@mdx-js/react'

const components = {
  CustomButton: CustomButton,
}

const Playground = () => {
  return (
    <MDXProvider components={components}>
      <div style={{ padding: '20px', maxWidth: '800px', margin: '0 auto' }}>
        <MyMdxPage />
      </div>
    </MDXProvider>
  )
}

export default Playground
