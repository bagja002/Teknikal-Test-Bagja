-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Mar 19, 2023 at 02:14 PM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kstyle_dev`
--

-- --------------------------------------------------------

--
-- Table structure for table `like_reviews`
--

CREATE TABLE `like_reviews` (
  `id_like` bigint(20) NOT NULL,
  `id_review` bigint(20) DEFAULT NULL,
  `id_member` bigint(20) DEFAULT NULL,
  `id_product` bigint(20) DEFAULT NULL,
  `is_like` int(100) NOT NULL COMMENT '1 = like, 0 = unlike'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `members`
--

CREATE TABLE `members` (
  `id_member` bigint(20) NOT NULL,
  `username` longtext DEFAULT NULL,
  `gender` longtext DEFAULT NULL COMMENT '0=wanita, 1 = pria',
  `skintype` longtext DEFAULT NULL,
  `skincolor` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `members`
--

INSERT INTO `members` (`id_member`, `username`, `gender`, `skintype`, `skincolor`) VALUES
(4, 'barjafaskan99@gmail.com', '1', 'dry', 'White'),
(5, 'barjafaskan20@gmail.com', '0', 'minyak', 'White');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id_product` bigint(20) NOT NULL,
  `name_product` longtext DEFAULT NULL,
  `price` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id_product`, `name_product`, `price`) VALUES
(1, 'Elvicto Men 80 Ml', 80000),
(2, 'Evict Serum Berjerawat ', 120000);

-- --------------------------------------------------------

--
-- Table structure for table `review_products`
--

CREATE TABLE `review_products` (
  `id_review` bigint(20) NOT NULL,
  `id_product` bigint(20) DEFAULT NULL,
  `id_member` bigint(20) DEFAULT NULL,
  `decs_review` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `like_reviews`
--
ALTER TABLE `like_reviews`
  ADD PRIMARY KEY (`id_like`),
  ADD KEY `id_review` (`id_review`);

--
-- Indexes for table `members`
--
ALTER TABLE `members`
  ADD PRIMARY KEY (`id_member`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id_product`);

--
-- Indexes for table `review_products`
--
ALTER TABLE `review_products`
  ADD PRIMARY KEY (`id_review`),
  ADD KEY `id_product` (`id_product`),
  ADD KEY `id_member` (`id_member`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `like_reviews`
--
ALTER TABLE `like_reviews`
  MODIFY `id_like` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `members`
--
ALTER TABLE `members`
  MODIFY `id_member` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id_product` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `review_products`
--
ALTER TABLE `review_products`
  MODIFY `id_review` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `like_reviews`
--
ALTER TABLE `like_reviews`
  ADD CONSTRAINT `like_reviews_ibfk_1` FOREIGN KEY (`id_review`) REFERENCES `review_products` (`id_review`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `review_products`
--
ALTER TABLE `review_products`
  ADD CONSTRAINT `review_products_ibfk_1` FOREIGN KEY (`id_member`) REFERENCES `members` (`id_member`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `review_products_ibfk_2` FOREIGN KEY (`id_product`) REFERENCES `products` (`id_product`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
