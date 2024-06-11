import './header.sass'
import { useRef } from 'react';

export default function Header() {
  const navRef = useRef<HTMLButtonElement>(null!);

  function toggleMenu() {
    navRef.current.classList.toggle('toggle_nav_menu');
  }
  
  return (
    <header>
      <div className='logo_container'>
        <h1>essentials</h1>
        <button className='menu_toggler' onClick={toggleMenu}>
          <hr />
          <hr />
          <hr />
        </button>
      </div>
      <nav ref={navRef}>
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
