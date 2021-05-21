-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 21 Bulan Mei 2021 pada 06.14
-- Versi server: 10.4.18-MariaDB
-- Versi PHP: 7.4.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `perpustakaan`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `author`
--

CREATE TABLE `author` (
  `id_author` int(11) NOT NULL,
  `nama_author` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `author`
--

INSERT INTO `author` (`id_author`, `nama_author`) VALUES
(1, 'Muhammad Hilmi Ramadhan'),
(2, 'Mr. Saul Muller DDS'),
(3, 'Virgie Paucek I'),
(4, 'Anastacio Cremin'),
(5, 'Cody Eichmann'),
(6, 'Stanford Hirthe'),
(7, 'Efrain Ullrich'),
(8, 'Dr. Carmel Bergstrom IV'),
(9, 'Prof. Rosendo Price I'),
(10, 'Demarco Carroll'),
(11, 'Chadrick Howell'),
(12, 'Prof. Jerry Labadie'),
(13, 'Mr. Columbus Wisoky'),
(14, 'Jody Hodkiewicz'),
(15, 'Carmen Bechtelar II'),
(16, 'Alfreda Kilback'),
(17, 'Prof. Vallie Fay PhD'),
(18, 'Caroline Mills'),
(19, 'Jules Ullrich'),
(20, 'Nadia Greenholt'),
(21, 'R.V. Hari Ginardi'),
(22, 'Baskoro Adi Pratomo'),
(26, 'Rizka Wakhidatus Shalikah');

-- --------------------------------------------------------

--
-- Struktur dari tabel `buku`
--

CREATE TABLE `buku` (
  `id_buku` int(11) NOT NULL,
  `id_author` int(11) NOT NULL,
  `nama_buku` varchar(100) DEFAULT NULL,
  `cover_book` varchar(200) DEFAULT NULL,
  `amountTotal` int(11) DEFAULT NULL,
  `amountOut` int(11) DEFAULT NULL,
  `amountRemaining` int(11) DEFAULT NULL,
  `klasifikasi` varchar(100) DEFAULT NULL,
  `penerbit` varchar(250) NOT NULL,
  `tahun` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `buku`
--

INSERT INTO `buku` (`id_buku`, `id_author`, `nama_buku`, `cover_book`, `amountTotal`, `amountOut`, `amountRemaining`, `klasifikasi`, `penerbit`, `tahun`) VALUES
(1, 8, 'Data Structure and Program Design in C', '1608702326_strukdat.jpg', 30, NULL, NULL, 'Pemrograman', 'Pearson', 1984),
(2, 7, 'Shopping with Dad', '1608720534_Shopping_with_Dad.jpg', 8, NULL, NULL, 'Children', 'Barefoot Books', 2010),
(3, 1, 'Teen Titans Go!', '1608723337_TeenTitansGo!.jpg', 30, NULL, NULL, 'Children', 'DC Comics', 2004),
(4, 1, 'The Very Hungry Caterpillar', '1608723430_TheVeryHungryCaterpillar.jpg', 12, NULL, NULL, 'Children', 'World Publishing Company', 1969),
(5, 3, 'The Lion, the Witch and the Wardrobe', '1608723509_The_Lion.jpg', 43, NULL, NULL, 'Children', 'C. S. Lewis', 1950),
(6, 20, 'Alices Adventures in Wonderland', '1608723712_Alices.jpg', 3, NULL, NULL, 'Children', 'Macmilian Publishers', 2016),
(7, 8, 'Think and Grow Rich', '1608727357_ThinkandGrowRich.jpg', 35, NULL, NULL, 'Business', 'The Ralston Society', 1937),
(8, 1, 'The 7 Habits of Highly Effective People', '1608727562_The7Habits.jpg', 20, NULL, NULL, 'Business', 'Free Press', 1989),
(9, 3, 'Thinking, Fast and Slow', '1608727767_ThinkingFastandSlow.jpg', 33, NULL, NULL, 'Business', 'Farrar, Straus, and Giroux', 2011),
(10, 18, 'The Lean Startup', '1608728026_TheLeanStartup.jpg', 9, NULL, NULL, 'Business', 'Eric Ries', 2011),
(11, 20, 'Deep Work', '1608728093_DeepWork.jpg', 37, NULL, NULL, 'Business', 'Grand Central Publishing', 2016),
(12, 5, 'Foundations of Bilingual Education and Bilingualism', '1608728724_Foundation.jpg', 31, NULL, NULL, 'Education', 'Colin Baker', 2009),
(13, 16, 'Physical Education: Essential Issues', '1608728845_Physical.jpg', 21, NULL, NULL, 'Education', 'a SAGE Publications Comp.', 2005),
(14, 15, 'WANGSIT (PAWANG SOAL SULIT) HOTS UTBK SBMPTN SAINTEK 2021', '1608728916_Wangsit.jpg', 8, NULL, NULL, 'Education', 'TIM TENTOR MASTER', 2009),
(15, 12, 'Buku Pintar Matematika SMP Untuk Kelas 1, 2 dan 3', '1608729303_BukuPintarMTK.jpg', 25, NULL, NULL, 'Education', 'Wahyu medfo', 2015),
(16, 1, 'Matematika SMP Kelas 8', '1608729492_Mtksmp-compressed.jpg', 23, NULL, NULL, 'Education', 'Gramedia', 2016),
(17, 5, 'Technology and Values: Essential Readings', '1608729842_TechnologyandValues.jpg', 15, NULL, NULL, 'Technology', 'Willey Blackwell', 2004),
(18, 14, 'Super Food Technology', '1608729914_SuperFood.jpg', 43, NULL, NULL, 'Technology', 'Marlen Sunyoto', 2010),
(19, 8, 'Artificial Intelligence for Humans', '1608730023_AIforHumans.jpg', 31, NULL, NULL, 'Technology', 'Jon Gabriel', 2007),
(20, 9, 'Agile Machine Learning', '1608730065_AgileML.jpg', 12, NULL, NULL, 'Technology', 'Apress', 2011),
(21, 1, 'The Essential PIC18 Microcontroller', '1608730226_PIC18.jpg', 41, NULL, NULL, 'Technology', 'Springer', 2009);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `username`, `password`) VALUES
(1, 'mhilmi', '$2a$10$4u0OZXEfSPoaUCp/epabnOYJ/jsJw9t2ar5SjwPqAh16fp8N/3cRi'),
(2, 'putri', '$2a$10$u1FY8CMLhSbNO2oYMK407uliDAjfP7NWAZt9gLGnpGPJWifrcmcJu');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `author`
--
ALTER TABLE `author`
  ADD PRIMARY KEY (`id_author`);

--
-- Indeks untuk tabel `buku`
--
ALTER TABLE `buku`
  ADD PRIMARY KEY (`id_buku`),
  ADD KEY `id_author` (`id_author`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `author`
--
ALTER TABLE `author`
  MODIFY `id_author` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;

--
-- AUTO_INCREMENT untuk tabel `buku`
--
ALTER TABLE `buku`
  MODIFY `id_buku` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `buku`
--
ALTER TABLE `buku`
  ADD CONSTRAINT `buku_ibfk_1` FOREIGN KEY (`id_author`) REFERENCES `author` (`id_author`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
