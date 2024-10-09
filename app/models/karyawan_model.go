package models

type Karyawan struct {
	KaryawanId      int     `json:"karyawan_id" gorm:"primaryKey"`
	KaryawanNama    *string `json:"karyawan_nama"`
	KaryawanFoto    *string `json:"karyawan_foto"`
	KaryawanJabatan *string `json:"karyawan_jabatan"`
	CabangId        int     `json:"cabang_id"`
	// KaryawanTanggalMasukKerja *string    `json:"karyawan_tanggal_masuk_kerja"`
	// KaryawanStatus            *string    `json:"karyawan_status"`
	// KaryawanToken             *string    `json:"karyawan_token"`
	// LockLocation              *string    `json:"lock_location"`
	// JabatanId                 *string    `json:"jabatan_id"`
	// AllowLogin                *string    `json:"allow_login"`
	// Role                      *string    `json:"role"`
	// CreatedAt                 *time.Time `json:"created_at"`

	// karyawan_username
	// karyawan_password
	// nik
	// karyawan_nama
	// karyawan_foto
	// karyawan_jabatan
	// agama
	// karyawan_tempat_lahir
	// karyawan_tanggal_lahir
	// karyawan_email
	// karyawan_no_hp
	// karyawan_alamat
	// karyawan_kota
	// karyawan_provinsi
	// karyawan_tanggal_masuk_kerja
	// karyawan_pendidikan_terakhir
	// status_pernikahan
	// status_pegawai
	// karyawan_status
	// karyawan_token
	// karyawan_handphone_code
	// karyawan_device_name
	// karyawan_reg_id
	// lock_location
	// bank_account_name
	// bank_name
	// bank_account_number
	// bank_branch
	// created_at
	// updated_at
	// jumlah_tanggungan
	// gaji_pokok
	// tunjangan_konsumsi
	// tunjangan_jabatan
	// tunjangan_kehadiran
	// tunjangan_luarkota
	// tunjangan_transport
	// tunjangan_kost
	// tunjangan_penampilan
	// bpjs
	// jumlah_cuti
	// jabatan_id
	// allow_login
	// absensi_count
	// late_count
	// sync_at
	// cabang_id
	// role
	// deleted_at
	// parent_id
}
