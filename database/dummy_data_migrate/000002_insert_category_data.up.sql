USE karma_mvc_foodopiaDB;
-- Insert Categories
INSERT INTO category (cat_name, cat_description) VALUES
-- Regional Categories
('South Indian', 'Traditional dishes from South India including dosas, idlis, sambars and coconut-based curries'),
('North Indian', 'Rich and flavorful dishes from North India with creamy gravies, rotis and tandoor items'),
('Gujarati', 'Sweet and savory vegetarian dishes from Gujarat with unique flavor combinations'),
('Punjabi', 'Hearty and robust dishes from Punjab with rich gravies and bread varieties'),
('Bengali', 'Traditional Bengali vegetarian dishes with subtle flavors and fish-free preparations'),
('Rajasthani', 'Royal cuisine from Rajasthan with rich gravies, dal preparations and traditional breads'),

-- Taste Categories
('Spicy', 'Hot and fiery dishes with intense heat levels and bold flavors'),
('Sweet', 'Dishes with natural or added sweetness, including desserts and sweet preparations'),
('Tangy', 'Dishes with acidic, sour or citrusy flavors that stimulate the palate'),
('Mild', 'Gentle flavors suitable for all age groups with minimal spice levels'),

-- Food Type Categories
('Light & Healthy', 'Nutritious, low-calorie dishes perfect for health-conscious diners'),
('Comfort Food', 'Hearty, satisfying dishes that provide emotional and physical comfort'),
('Street Food', 'Popular roadside snacks and quick bites with authentic Indian flavors'),
('Desserts', 'Traditional Indian sweets and desserts to end meals on a sweet note'),
('Breakfast', 'Morning meals and dishes traditionally consumed to start the day'),
('Main Course', 'Substantial dishes that form the centerpiece of a complete meal'),

-- Cooking Style Categories
('Fried', 'Deep-fried or pan-fried dishes with crispy textures and rich flavors'),
('Steamed', 'Healthy steamed preparations that retain natural flavors and nutrients'),
('Grilled', 'Tandoor or grill-cooked items with smoky flavors and charred textures'),
('Curry', 'Gravy-based dishes with complex spice blends and rich sauces'),

-- Special Diet Categories
('Vegan', 'Plant-based dishes without any dairy or animal products'),
('Gluten-Free', 'Dishes made without wheat, suitable for gluten-sensitive individuals'),
('Jain', 'Vegetarian dishes prepared without onion, garlic, and root vegetables');
