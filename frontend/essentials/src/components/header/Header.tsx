import './header.sass'
import { useEffect, useRef, useState } from 'react'

export default function Header() {
  const [isOpen, setIsOpen] = useState(false);
  const headerRef = useRef<HTMLElement>(null!);

  const toggleMenu = (event: MouseEvent) => {
    if (!event.composedPath().includes(headerRef.current) && isOpen) {
      setIsOpen(false);
    }
  }

  useEffect(() => {
    document.addEventListener('click', toggleMenu);
    return () => {
      document.removeEventListener('click', toggleMenu);
    }
  })

  return (
    <header className='header' ref={headerRef}>
      <div className='logo_container'>
        <h1>essentials</h1>
        <button 
          className='menu_toggler'
          onClick={() => setIsOpen(prev => !prev)}>
          <hr />
          <hr />
          <hr />
        </button>
      </div>
      <nav className={isOpen ? 'toggle_nav_menu' : ''}>
        <ul>
          <li><a href="#">Home</a></li>
          <li><a href="#">Shop</a></li>
          <li><a href="#">About</a></li>
          <li><a href="#">Contact us</a></li>
        </ul>
      </nav>
    </header>
  )
}
