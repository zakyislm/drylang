// 04_control_flow.y - If, Elif, El, and Loops

score 85

// Conditional
if score >= 90 {
    pt "Grade: A"
} elif score >= 80 {
    pt "Grade: B"
} el {
    pt "Grade: C"
}

// Loop (lp) - repeats N times, exposing 'i' as index
pt "Counting to 3:"
lp 3 {
    pt "Iteration: ${i}"
}
