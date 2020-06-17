# paletteGenerator
A simply go module that generates a color palette according to a perceptual model of color vision

This random palette generator is based on a perceptual model of color vision, which says that, while 
the cones in our eyes are sensitive to specific frequency bands, our visual cortex codifies colors
based on where they lie among three different axes, at the ends of which are certain color components.
Specifically this program uses a Green-Magenta, Red-Cyan and Blue-Yellow choice of coordinates.
Thus we generate colors that are as distant as possible in this choice of coordinates, specifically generating a 
cubical crystal, the nodes of which are our colors. We skip the nodes on the diagonal, 
as they are less visibile according to our model, and eventually convert back to an RGB color model for ease of use.
