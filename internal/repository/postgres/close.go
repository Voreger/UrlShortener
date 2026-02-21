package postgres

// Close database connection pool
func (r *Repository) Close() {
	r.pool.Close()
}
