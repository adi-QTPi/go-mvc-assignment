
-- Insert Items with semantic relationships to categories
INSERT INTO item (item_name, cook_time_min, price, cat_id, subcat_id) VALUES
-- South Indian Items
('Masala Dosa', 25, 120, 1, 15), -- South Indian + Breakfast
('Idli Sambar', 20, 80, 1, 18), -- South Indian + Steamed
('Rava Upma', 15, 70, 1, 11), -- South Indian + Light & Healthy
('Medu Vada', 20, 90, 1, 17), -- South Indian + Fried
('Coconut Rice', 30, 110, 1, 16), -- South Indian + Main Course
('Filter Coffee', 5, 40, 1, 10), -- South Indian + Mild
('Rasam', 25, 60, 1, 9), -- South Indian + Tangy
('Pongal', 30, 100, 1, 12), -- South Indian + Comfort Food

-- North Indian Items
('Paneer Butter Masala', 35, 180, 2, 20), -- North Indian + Curry
('Chole Bhature', 30, 150, 2, 13), -- North Indian + Street Food
('Rajma Rice', 40, 140, 2, 12), -- North Indian + Comfort Food
('Aloo Paratha', 25, 100, 2, 15), -- North Indian + Breakfast
('Dal Makhani', 45, 160, 2, 20), -- North Indian + Curry
('Butter Naan', 15, 50, 2, 16), -- North Indian + Main Course
('Tandoori Roti', 12, 35, 2, 19), -- North Indian + Grilled

-- Gujarati Items
('Dhokla', 20, 80, 3, 18), -- Gujarati + Steamed
('Khandvi', 25, 90, 3, 11), -- Gujarati + Light & Healthy
('Gujarati Thali', 15, 220, 3, 16), -- Gujarati + Main Course
('Fafda Jalebi', 20, 110, 3, 8), -- Gujarati + Sweet
('Handvo', 35, 120, 3, 17), -- Gujarati + Fried
('Undhiyu', 50, 180, 3, 7), -- Gujarati + Spicy

-- Punjabi Items
('Sarson Ka Saag', 45, 160, 4, 20), -- Punjabi + Curry
('Makki Ki Roti', 20, 60, 4, 16), -- Punjabi + Main Course
('Punjabi Kadhi', 35, 140, 4, 9), -- Punjabi + Tangy
('Stuffed Kulcha', 25, 120, 4, 17), -- Punjabi + Fried
('Amritsari Kulcha', 30, 130, 4, 13), -- Punjabi + Street Food

-- Bengali Items
('Aloo Posto', 30, 120, 5, 10), -- Bengali + Mild
('Bengali Dal', 25, 100, 5, 20), -- Bengali + Curry
('Cholar Dal', 35, 130, 5, 8), -- Bengali + Sweet
('Aloo Bhaja', 15, 70, 5, 17), -- Bengali + Fried

-- Rajasthani Items
('Dal Baati Churma', 60, 200, 6, 12), -- Rajasthani + Comfort Food
('Gatte Ki Sabji', 40, 150, 6, 7), -- Rajasthani + Spicy
('Bajre Ki Roti', 20, 45, 6, 16), -- Rajasthani + Main Course
('Ker Sangri', 35, 140, 6, 9), -- Rajasthani + Tangy

-- Spicy Items
('Mirchi Bada', 20, 85, 7, 17), -- Spicy + Fried
('Spicy Paneer Tikka', 25, 180, 7, 19), -- Spicy + Grilled
('Achari Aloo', 30, 110, 7, 20), -- Spicy + Curry

-- Sweet Items
('Gulab Jamun', 15, 80, 8, 14), -- Sweet + Desserts
('Rasgulla', 10, 70, 8, 14), -- Sweet + Desserts
('Kheer', 30, 90, 8, 14), -- Sweet + Desserts
('Jalebi', 15, 60, 8, 17), -- Sweet + Fried

-- Tangy Items
('Pani Puri', 10, 50, 9, 13), -- Tangy + Street Food
('Bhel Puri', 8, 60, 9, 13), -- Tangy + Street Food
('Aam Panna', 5, 45, 9, 11), -- Tangy + Light & Healthy

-- Mild Items
('Plain Rice', 20, 60, 10, 16), -- Mild + Main Course
('Khichdi', 25, 80, 10, 12), -- Mild + Comfort Food
('Buttermilk', 3, 30, 10, 11), -- Mild + Light & Healthy

-- Light & Healthy Items
('Sprouts Salad', 10, 70, 11, 21), -- Light & Healthy + Vegan
('Vegetable Soup', 15, 65, 11, 18), -- Light & Healthy + Steamed
('Fruit Chaat', 8, 80, 11, 9), -- Light & Healthy + Tangy
('Green Salad', 5, 50, 11, 21), -- Light & Healthy + Vegan

-- Comfort Food Items
('Poha', 15, 60, 12, 15), -- Comfort Food + Breakfast
('Maggi Noodles', 8, 55, 12, 7), -- Comfort Food + Spicy
('Aloo Parantha', 25, 95, 12, 17), -- Comfort Food + Fried

-- Street Food Items
('Vada Pav', 12, 40, 13, 17), -- Street Food + Fried
('Pav Bhaji', 20, 90, 13, 7), -- Street Food + Spicy
('Sev Puri', 8, 55, 13, 9), -- Street Food + Tangy
('Dahi Puri', 10, 65, 13, 8), -- Street Food + Sweet

-- Dessert Items
('Kulfi', 5, 60, 14, 8), -- Desserts + Sweet
('Ras Malai', 12, 90, 14, 8), -- Desserts + Sweet
('Gajar Halwa', 35, 100, 14, 8), -- Desserts + Sweet

-- Breakfast Items
('Aloo Poha', 15, 70, 15, 10), -- Breakfast + Mild
('Upma', 12, 60, 15, 1), -- Breakfast + South Indian
('Paratha with Curd', 20, 85, 15, 10), -- Breakfast + Mild

-- Fried Items
('Samosa', 18, 45, 17, 13), -- Fried + Street Food
('Pakora', 15, 55, 17, 7), -- Fried + Spicy
('Aloo Tikki', 20, 70, 17, 13), -- Fried + Street Food

-- Steamed Items
('Vegetable Momos', 20, 120, 18, 11), -- Steamed + Light & Healthy
('Kothimbir Vadi', 25, 80, 18, 11), -- Steamed + Light & Healthy

-- Grilled Items
('Paneer Tikka', 25, 160, 19, 10), -- Grilled + Mild
('Grilled Vegetables', 20, 140, 19, 11), -- Grilled + Light & Healthy

-- Curry Items
('Mixed Vegetable Curry', 35, 130, 20, 16), -- Curry + Main Course
('Palak Paneer', 30, 150, 20, 11), -- Curry + Light & Healthy

-- Vegan Items
('Baingan Bharta', 40, 120, 21, 20), -- Vegan + Curry
('Aloo Jeera', 20, 90, 21, 10), -- Vegan + Mild
('Chana Masala', 35, 110, 21, 7), -- Vegan + Spicy

-- Jain Items
('Jain Pav Bhaji', 25, 95, 23, 13), -- Jain + Street Food
('Jain Pizza', 20, 180, 23, 16), -- Jain + Main Course
('Jain Fried Rice', 25, 130, 23, 17); -- Jain + Fried