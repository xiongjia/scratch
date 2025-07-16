import { Button } from 'antd'

interface CustomButtonProps {
  children: React.ReactNode
  onClick?: () => void
  variant?: 'primary' | 'secondary'
}

const CustomButton: React.FC<CustomButtonProps> = ({
  children,
  onClick,
  variant = 'primary',
}) => {
  const baseStyle = 'px-4 py-2 rounded font-semibold transition duration-200'
  const primaryStyle = 'bg-blue-500 text-white hover:bg-blue-600'
  const secondaryStyle = 'bg-gray-200 text-gray-800 hover:bg-gray-300'

  return (
    <Button
      onClick={onClick}
      style={{
        margin: '5px',
        border: 'none',
        cursor: 'pointer',
        ...(variant === 'primary'
          ? { backgroundColor: '#007bff', color: 'white' }
          : { backgroundColor: '#e0e0e0', color: '#333' }),
      }}
      className={`${baseStyle} ${variant === 'primary' ? primaryStyle : secondaryStyle}`}
    >
      {children}
    </Button>
  )
}

export default CustomButton
