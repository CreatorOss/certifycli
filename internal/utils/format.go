package utils

import (
	"fmt"
	"strings"
	"time"
)

// FormatTime formats time for pretty output
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FormatDuration formats duration human-readably
func FormatDuration(d time.Duration) string {
	days := d / (24 * time.Hour)
	hours := (d % (24 * time.Hour)) / time.Hour
	minutes := (d % time.Hour) / time.Minute

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}

// BoxedMessage creates a boxed message for emphasis
func BoxedMessage(title, message string) string {
	width := 50
	lines := strings.Split(message, "\n")
	
	var maxLen int
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	if maxLen < len(title) {
		maxLen = len(title)
	}
	maxLen += 4

	if maxLen > width {
		maxLen = width
	}

	var box strings.Builder
	box.WriteString("╭" + strings.Repeat("─", maxLen-2) + "╮\n")
	box.WriteString(fmt.Sprintf("│ %-*s │\n", maxLen-3, title))
	box.WriteString("├" + strings.Repeat("─", maxLen-2) + "┤\n")
	
	for _, line := range lines {
		box.WriteString(fmt.Sprintf("│ %-*s │\n", maxLen-3, line))
	}
	
	box.WriteString("╰" + strings.Repeat("─", maxLen-2) + "╯\n")
	return box.String()
}

// Color constants
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorReset  = "\033[0m"
	ColorBold   = "\033[1m"
)

// Colorize returns colored string
func Colorize(color, text string) string {
	return color + text + ColorReset
}

// Success returns green success message
func Success(text string) string {
	return Colorize(ColorGreen, "✅ "+text)
}

// Error returns red error message
func Error(text string) string {
	return Colorize(ColorRed, "❌ "+text)
}

// Warning returns yellow warning message
func Warning(text string) string {
	return Colorize(ColorYellow, "⚠️  "+text)
}

// Info returns blue info message
func Info(text string) string {
	return Colorize(ColorBlue, "ℹ️  "+text)
}

// Bold returns bold text
func Bold(text string) string {
	return Colorize(ColorBold, text)
}

// FormatTable creates a simple table
func FormatTable(headers []string, rows [][]string) string {
	if len(rows) == 0 {
		return "No data to display"
	}

	// Calculate column widths
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = len(header)
	}

	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	var table strings.Builder

	// Header
	table.WriteString("┌")
	for i, width := range colWidths {
		table.WriteString(strings.Repeat("─", width+2))
		if i < len(colWidths)-1 {
			table.WriteString("┬")
		}
	}
	table.WriteString("┐\n")

	// Header row
	table.WriteString("│")
	for i, header := range headers {
		table.WriteString(fmt.Sprintf(" %-*s │", colWidths[i], header))
	}
	table.WriteString("\n")

	// Header separator
	table.WriteString("├")
	for i, width := range colWidths {
		table.WriteString(strings.Repeat("─", width+2))
		if i < len(colWidths)-1 {
			table.WriteString("┼")
		}
	}
	table.WriteString("┤\n")

	// Data rows
	for _, row := range rows {
		table.WriteString("│")
		for i, cell := range row {
			if i < len(colWidths) {
				table.WriteString(fmt.Sprintf(" %-*s │", colWidths[i], cell))
			}
		}
		table.WriteString("\n")
	}

	// Footer
	table.WriteString("└")
	for i, width := range colWidths {
		table.WriteString(strings.Repeat("─", width+2))
		if i < len(colWidths)-1 {
			table.WriteString("┴")
		}
	}
	table.WriteString("┘\n")

	return table.String()
}

// ProgressBar creates a simple progress bar
func ProgressBar(current, total int, width int) string {
	if total == 0 {
		return "[" + strings.Repeat("─", width) + "] 0%"
	}

	percentage := float64(current) / float64(total)
	filled := int(percentage * float64(width))
	
	bar := "["
	bar += strings.Repeat("█", filled)
	bar += strings.Repeat("─", width-filled)
	bar += fmt.Sprintf("] %.1f%% (%d/%d)", percentage*100, current, total)
	
	return bar
}