package certstore

// certs to store

const (
	pem__creator                   = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNJakNDQWNpZ0F3SUJBZ0lRSDlSVXpKWEEwKzBEeFh1d2hMZS9jVEFLQmdncWhrak9QUVFEQWpCNU1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RWNNQm9HQTFVRUNoTVRaWFp5WVhwNmMyMXJMbkpoYVd4ekxuSjZaREVmTUIwR0ExVUVBeE1XClkyRXVaWFp5WVhwNmMyMXJMbkpoYVd4ekxuSjZaREFlRncweU1UQTNNVFF3T0RFM01EQmFGdzB6TVRBM01USXcKT0RFM01EQmFNRjR4Q3pBSkJnTlZCQVlUQWxWVE1STXdFUVlEVlFRSUV3cERZV3hwWm05eWJtbGhNUll3RkFZRApWUVFIRXcxVFlXNGdSbkpoYm1OcGMyTnZNU0l3SUFZRFZRUUREQmxCWkcxcGJrQmxkbkpoZW5wemJXc3VjbUZwCmJITXVjbnBrTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFcDIydDZJWEl1cVZkWW0vWnVzL3YKM3d6Y3A4RlRkVGFDQlJPRDlzeGdDbDZtNHY2ekVoU0dmOEZ5UmhuS2s0cHo5U3NJUVphQ0RHaVYzdUpVQnlITgpTYU5OTUVzd0RnWURWUjBQQVFIL0JBUURBZ2VBTUF3R0ExVWRFd0VCL3dRQ01BQXdLd1lEVlIwakJDUXdJb0FnCmQ1SGEvU3Vab3J4Q2FkeEo3TVQvNTJxZFhZK3pobm8xNTgvWHFDaFRQMjR3Q2dZSUtvWkl6ajBFQXdJRFNBQXcKUlFJaEFNWkhnMHBkNHhEWHpUUUcrTXFZSnVtZDd2QjFwN29wYmYvR2ZzSnREVnZvQWlCWXIrajRLNmlha2NFNgo3Sy9ickNMbm9DUXlqQlpOM1VCdnJEZDI2bDdrbWc9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
	pem__peer0_rzd_rails_rzd       = "CgZSWkRNU1ASigYtLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJQ0VUQ0NBYmVnQXdJQkFnSVJBTlozb05NNHFrRzV2WE5WdWtsT0Ruc3dDZ1lJS29aSXpqMEVBd0l3YlRFTApNQWtHQTFVRUJoTUNWVk14RXpBUkJnTlZCQWdUQ2tOaGJHbG1iM0p1YVdFeEZqQVVCZ05WQkFjVERWTmhiaUJHCmNtRnVZMmx6WTI4eEZqQVVCZ05WQkFvVERYSjZaQzV5WVdsc2N5NXllbVF4R1RBWEJnTlZCQU1URUdOaExuSjYKWkM1eVlXbHNjeTV5ZW1Rd0hoY05NakV3TnpFME1EZ3hOekF3V2hjTk16RXdOekV5TURneE56QXdXakJZTVFzdwpDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTlUyRnVJRVp5CllXNWphWE5qYnpFY01Cb0dBMVVFQXhNVGNHVmxjakF1Y25wa0xuSmhhV3h6TG5KNlpEQlpNQk1HQnlxR1NNNDkKQWdFR0NDcUdTTTQ5QXdFSEEwSUFCTFluTUo5N2t4L0cxOENLUGhRMUpmSWFNenpBaEUrWXhGbmFuMTRSVFRxQgpxby9LaXpjSlJZTjVxOVNkK0sxNnVQUG56VFF2K0VxeVk5YWErM2J4RHhPalRUQkxNQTRHQTFVZER3RUIvd1FFCkF3SUhnREFNQmdOVkhSTUJBZjhFQWpBQU1Dc0dBMVVkSXdRa01DS0FJQnh0YTNycFZCcVhpbmpsQ0EyWFowZ1cKRDNTREp5dGgwak5Dbi9HODd0eEpNQW9HQ0NxR1NNNDlCQU1DQTBnQU1FVUNJUUNBVUUyUmRsckpsWWR6S20xLwpnYm5jMklob0wvUXluU0paWHl1T1JtZXRPZ0lnQitYWnNIZ1lSUlBFeXJvczFocnY3dDhsVDdtUWIyaWlPeDJoClBJQTNCS1U9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
	pem__peer0_hmk_rails_rzd       = "CgZITUtNU1AShgYtLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJQ0R6Q0NBYmFnQXdJQkFnSVFSMVh4K2NTaEJ2YXZLSlRUMnZyTXRqQUtCZ2dxaGtqT1BRUURBakJ0TVFzdwpDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTlUyRnVJRVp5CllXNWphWE5qYnpFV01CUUdBMVVFQ2hNTmFHMXJMbkpoYVd4ekxuSjZaREVaTUJjR0ExVUVBeE1RWTJFdWFHMXIKTG5KaGFXeHpMbko2WkRBZUZ3MHlNVEEzTVRRd09ERTNNREJhRncwek1UQTNNVEl3T0RFM01EQmFNRmd4Q3pBSgpCZ05WQkFZVEFsVlRNUk13RVFZRFZRUUlFd3BEWVd4cFptOXlibWxoTVJZd0ZBWURWUVFIRXcxVFlXNGdSbkpoCmJtTnBjMk52TVJ3d0dnWURWUVFERXhOd1pXVnlNQzVvYldzdWNtRnBiSE11Y25wa01Ga3dFd1lIS29aSXpqMEMKQVFZSUtvWkl6ajBEQVFjRFFnQUVyUm0zQm53REpCYzB0bk1PZ2F4OTd1ZWY1MjFDRnpHbFcxSG5ZNTdlclpzZAp3LzJpSExBdmkvb0pLNHRZUS9pd2hYTjBEMTZ5Z0g3eWEySkhqYTVjTGFOTk1Fc3dEZ1lEVlIwUEFRSC9CQVFECkFnZUFNQXdHQTFVZEV3RUIvd1FDTUFBd0t3WURWUjBqQkNRd0lvQWdRUVdOVUx5RTdDNTFOTVRJTGVPL1h6WkoKL2pzTllnMld1UjB5MERKV1loOHdDZ1lJS29aSXpqMEVBd0lEUndBd1JBSWdZL00wYmpSZ1JqaHFWSTJGeXB5cwpBSHNwYnI3T3cvS1lQckhNT3BaUlEyc0NJRldySW9NTGNuTDdoQk1oWDRHczdKdEU2RjZFcmtwbmp2ME16UHRKCmxuYksKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
	pem__peer0_evrazzsmk_rails_rzd = "CgxFVlJBWlpTTUtNU1ASngYtLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJQ0lUQ0NBY2lnQXdJQkFnSVFNNTlCbWNOeHVmNjZ3QldPUHhWY3pEQUtCZ2dxaGtqT1BRUURBakI1TVFzdwpDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTlUyRnVJRVp5CllXNWphWE5qYnpFY01Cb0dBMVVFQ2hNVFpYWnlZWHA2YzIxckxuSmhhV3h6TG5KNlpERWZNQjBHQTFVRUF4TVcKWTJFdVpYWnlZWHA2YzIxckxuSmhhV3h6TG5KNlpEQWVGdzB5TVRBM01UUXdPREUzTURCYUZ3MHpNVEEzTVRJdwpPREUzTURCYU1GNHhDekFKQmdOVkJBWVRBbFZUTVJNd0VRWURWUVFJRXdwRFlXeHBabTl5Ym1saE1SWXdGQVlEClZRUUhFdzFUWVc0Z1JuSmhibU5wYzJOdk1TSXdJQVlEVlFRREV4bHdaV1Z5TUM1bGRuSmhlbnB6YldzdWNtRnAKYkhNdWNucGtNRmt3RXdZSEtvWkl6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUVGZ0Q5cnZCd3VNREtCcFRDUmNYRwpFMUpXUG5FRVV5R3NIQ3RBb245RXQ5aGNRd1hLUnNLVytsdmsxMFRkT2FFT3h6d0tDa09ISHJneXQ2MkliNlRWCmhhTk5NRXN3RGdZRFZSMFBBUUgvQkFRREFnZUFNQXdHQTFVZEV3RUIvd1FDTUFBd0t3WURWUjBqQkNRd0lvQWcKZDVIYS9TdVpvcnhDYWR4SjdNVC81MnFkWFkremhubzE1OC9YcUNoVFAyNHdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnREU5TXU1SDFQby9USW5nUFZyeHFzMVVzWlRrRERtZzhEdkczTjlXR1g3NENJRytpdnNvdFFCOGVqUm9HCmwxVHRjSGJKNG1UY1FMRjF3UGFUdjFKR0NYSk8KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
)

var certs = []string{
	pem__creator,
	pem__peer0_rzd_rails_rzd,
	pem__peer0_hmk_rails_rzd,
	pem__peer0_evrazzsmk_rails_rzd,
}