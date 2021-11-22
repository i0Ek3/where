/*
 * Package where finds where the data offered. We offer three method
 * to get the data, which includes FromArgs(), FromFile() and FromStdin().
 *
 * FromArgs read args and checks if an arg is a file, if true, goto FromFile(),
 * if not, save the args.
 *
 * FromFile read filename..., if there have no filename offered, serach
 * default file locally, otherwise checks if filename... is a list,
 * if len(filename) == 1, then invoke processFile(), if len(filename) > 1,
 * open and read these files one by one.
 *
 * FromStdin just read input from os.Stdin, and then convert and return
 * the input.
 *
 * You can invoke one of these three methods to support the data easily,
 * instead of manage it manually.
 *
 */
 package where
